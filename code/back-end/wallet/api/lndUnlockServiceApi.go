package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
)

// GenSeed
//
//	 @Description: 是用于实例化新 lnd 实例的第一个方法。该方法允许调用者根据可选的口令生成新的加密种子。
//		如果提供了口令，则需要口令来解密密码种子，以显示内部钱包种子。
//		用户获得并验证密码种子后，应使用 InitWallet 方法提交新生成的种子，并创建钱包
//	 @return []string
func GenSeed() []string {
	const (
		//grpcHost = "202.79.173.41:10009"
		grpcHost = "202.79.173.41:10055"
	)
	tlsCertPath := filepath.Join(base.Configure("lnd2"), "tls.cert")
	// Load the TLS certificate
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		log.Printf("Failed to read cert file: %s", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	// 创建 WalletUnlocker 客户端
	client := lnrpc.NewWalletUnlockerClient(conn)
	passphrase := ""
	var aezeedPassphrase = []byte(passphrase)
	// Generating random seed entropy:
	seedEntropy := make([]byte, 16)
	_, err = rand.Read(seedEntropy)
	if err != nil {
		log.Printf("could not generate seed entropy: %v", err)
	}
	// 准备 WalletUnlocker 请求
	request := &lnrpc.GenSeedRequest{
		AezeedPassphrase: aezeedPassphrase,
		SeedEntropy:      seedEntropy,
	}
	// 调用 InitWallet gRPC 方法
	response, err := client.GenSeed(context.Background(), request)
	if err != nil {
		log.Printf("Error calling InitWallet: %v", err)
	}
	// 处理 gRPC 响应
	return response.CipherSeedMnemonic
}

// InitWallet
//
//	@Description: InitWallet 在 lnd 首次启动时使用，用于完全初始化守护进程及其内部钱包。
//	至少必须提供一个钱包密码。这将用于加密磁盘上的敏感资料。
//	在恢复情况下，用户还可以指定自己的密码和口令。如果设置了该密码，守护进程就会使用之前的状态来初始化其内部钱包。
//	或者，也可以使用 GenSeed RPC 来获取种子，然后将其提交给用户。经用户验证后，可将种子输入此 RPC，以提交新钱包
//	@param seed
//	@param password
//	@return bool
func InitWallet(seed [24]string, password string) bool {
	const (
		//grpcHost = "202.79.173.41:10009"
		grpcHost = "202.79.173.41:10055"
	)
	tlsCertPath := filepath.Join(base.Configure("lnd2"), "tls.cert")
	// Load the TLS certificate
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		log.Printf("Failed to read cert file: %s", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	// 创建 WalletUnlocker 客户端
	client := lnrpc.NewWalletUnlockerClient(conn)
	passphrase := ""
	request := &lnrpc.InitWalletRequest{
		WalletPassword:     []byte(password),
		CipherSeedMnemonic: seed[:],
		AezeedPassphrase:   []byte(passphrase),
	}
	response, err := client.InitWallet(context.Background(), request)
	if err != nil {
		log.Printf("Error calling InitWallet: %v", err)
	}
	d1 := response.AdminMacaroon
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	err = os.MkdirAll(newFilePath, os.ModePerm)
	if err != nil {
		log.Printf("Error calling InitWallet: %v", err)
	}
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	f, err := os.Create(macaroonPath)
	if err != nil {
		log.Printf("Error calling InitWallet: %v", err)
		return false
	}
	_, err = f.Write(d1)
	if err != nil {
		err := f.Close()
		if err != nil {
			log.Printf("f Close err: %v\n", err)
			return false
		}
		log.Printf("Error calling InitWallet: %v", err)
		return false
	}
	log.Println("successful")
	err = f.Close()
	if err != nil {
		log.Printf("Error calling InitWallet: %v", err)
		return false
	}
	return true
}

func UnlockWallet(password string) bool {
	const (
		grpcHost = "202.79.173.41:10009"
	)
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	// Load the TLS certificate
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		log.Printf("Failed to read cert file: %s", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	// 创建 WalletUnlocker 客户端
	client := lnrpc.NewWalletUnlockerClient(conn)
	request := &lnrpc.UnlockWalletRequest{
		WalletPassword: []byte(password),
	}
	_, err = client.UnlockWallet(context.Background(), request)
	if err != nil {
		log.Printf("did not connect: %v", err)
		return false
	}
	log.Println("unlockSuccess")
	return true
}

func ChangePassword(currentPassword, newPassword string) bool {
	const (
		grpcHost = "202.79.173.41:10009"
	)
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
		log.Printf("Failed to read cert file: %s", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	client := lnrpc.NewWalletUnlockerClient(conn)
	request := &lnrpc.ChangePasswordRequest{
		CurrentPassword: []byte(currentPassword),
		NewPassword:     []byte(newPassword),
	}
	response, err := client.ChangePassword(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ChangePassword err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}
