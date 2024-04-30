package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"os"
	"path/filepath"
	"strings"
)

// GenSeed
//
//	@Description: GenSeed is the first method that should be used to instantiate a new lnd instance.
//	This method allows a caller to generate a new aezeed cipher seed given an optional passphrase.
//	If provided, the passphrase will be necessary to decrypt the cipherseed to expose the internal wallet seed.
//	Once the cipherseed is obtained and verified by the user, the InitWallet method should be used to commit the newly generated seed, and create the wallet.
//	@return string
func GenSeed() string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	creds := NewTlsCert(tlsCertPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewWalletUnlockerClient(conn)
	//passphrase := ""
	//var aezeedPassphrase = []byte(passphrase)
	seedEntropy := make([]byte, 16)
	_, err = rand.Read(seedEntropy)
	if err != nil {
		fmt.Printf("%s could not generate seed entropy: %v\n", GetTimeNow(), err)
	}
	request := &lnrpc.GenSeedRequest{
		//AezeedPassphrase: aezeedPassphrase,
		//SeedEntropy:      seedEntropy,
	}
	response, err := client.GenSeed(context.Background(), request)
	if err != nil {
		fmt.Printf("%s Error calling GenSeed: %v\n", GetTimeNow(), err)
	}
	return strings.Join(response.CipherSeedMnemonic, ",")
}

// InitWallet
//
//	@Description:InitWallet is used when lnd is starting up for the first time to fully initialize the daemon and its internal wallet. At the very least a wallet password must be provided.
//	This will be used to encrypt sensitive material on disk.
//	In the case of a recovery scenario, the user can also specify their aezeed mnemonic and passphrase.
//	If set, then the daemon will use this prior state to initialize its internal wallet.
//	Alternatively, this can be used along with the GenSeed RPC to obtain a seed, then present it to the user.
//	Once it has been verified by the user, the seed can be fed into this RPC in order to commit the new wallet.
//	@return bool
func InitWallet(seed, password string) bool {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	creds := NewTlsCert(tlsCertPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)

	var (
		cipherSeedMnemonic      []string
		aezeedPass              []byte
		extendedRootKey         string
		extendedRootKeyBirthday uint64
		recoveryWindow          int32
	)

	client := lnrpc.NewWalletUnlockerClient(conn)
	//seedrequest := &lnrpc.GenSeedRequest{}
	//seedresponse, err := client.GenSeed(context.Background(), seedrequest)
	//cipherSeedMnemonic = seedresponse.CipherSeedMnemonic
	//
	//recoveryWindow = 2500
	cipherSeedMnemonic = strings.Split(seed, ",")
	request := &lnrpc.InitWalletRequest{
		WalletPassword:                     []byte(password),
		CipherSeedMnemonic:                 cipherSeedMnemonic,
		AezeedPassphrase:                   aezeedPass,
		RecoveryWindow:                     recoveryWindow,
		ChannelBackups:                     nil,
		StatelessInit:                      false,
		ExtendedMasterKey:                  extendedRootKey,
		ExtendedMasterKeyBirthdayTimestamp: extendedRootKeyBirthday,
	}
	response, err := client.InitWallet(context.Background(), request)
	if err != nil {
		fmt.Printf("%s Error calling InitWallet: %v\n", GetTimeNow(), err)
	}
	d1 := response.AdminMacaroon
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	err = os.MkdirAll(newFilePath, os.ModePerm)
	if err != nil {
		fmt.Printf("%s Error calling InitWallet: %v\n", GetTimeNow(), err)
	}
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	f, err := os.Create(macaroonPath)
	if err != nil {
		fmt.Printf("%s Error calling InitWallet: %v\n", GetTimeNow(), err)
		return false
	}
	_, err = f.Write(d1)
	if err != nil {
		err := f.Close()
		if err != nil {
			fmt.Printf("%s f Close err: %v\n", GetTimeNow(), err)
			return false
		}
		fmt.Printf("%s Error calling InitWallet: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s successful\n", GetTimeNow())
	err = f.Close()
	if err != nil {
		fmt.Printf("%s Error calling InitWallet: %v\n", GetTimeNow(), err)
		return false
	}
	return true
}

// UnlockWallet
//
//	@Description: UnlockWallet is used at startup of lnd to provide a password to unlock the wallet database.
//	@return bool
func UnlockWallet(password string) bool {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	creds := NewTlsCert(tlsCertPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewWalletUnlockerClient(conn)
	request := &lnrpc.UnlockWalletRequest{
		WalletPassword: []byte(password),
	}
	_, err = client.UnlockWallet(context.Background(), request)
	if err != nil {
		fmt.Printf("%s did not UnlockWallet: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s unlockSuccess\n", GetTimeNow())
	return true
}

// ChangePassword
//
//	@Description:ChangePassword changes the password of the encrypted wallet.
//	This will automatically unlock the wallet database if successful.
//	@return bool
func ChangePassword(currentPassword, newPassword string) bool {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	creds := NewTlsCert(tlsCertPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewWalletUnlockerClient(conn)
	request := &lnrpc.ChangePasswordRequest{
		CurrentPassword: []byte(currentPassword),
		NewPassword:     []byte(newPassword),
	}
	_, err = client.ChangePassword(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ChangePassword err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s ChangePassword Successfully\n", GetTimeNow())
	return true
}
