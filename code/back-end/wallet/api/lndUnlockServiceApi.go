package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
)

func GenSeed() [24]string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
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
	passphrase := ""
	var aezeedPassphrase = []byte(passphrase)
	seedEntropy := make([]byte, 16)
	_, err = rand.Read(seedEntropy)
	if err != nil {
		fmt.Printf("%s could not generate seed entropy: %v\n", GetTimeNow(), err)
	}
	request := &lnrpc.GenSeedRequest{
		AezeedPassphrase: aezeedPassphrase,
		SeedEntropy:      seedEntropy,
	}
	response, err := client.GenSeed(context.Background(), request)
	if err != nil {
		fmt.Printf("%s Error calling InitWallet: %v\n", GetTimeNow(), err)
	}
	return [24]string(response.CipherSeedMnemonic)
}

func InitWallet(seed [24]string, password string) bool {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
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
	passphrase := ""
	request := &lnrpc.InitWalletRequest{
		WalletPassword:     []byte(password),
		CipherSeedMnemonic: seed[:],
		AezeedPassphrase:   []byte(passphrase),
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

func UnlockWallet(password string) bool {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
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
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s unlockSuccess\n", GetTimeNow())
	return true
}

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
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
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
