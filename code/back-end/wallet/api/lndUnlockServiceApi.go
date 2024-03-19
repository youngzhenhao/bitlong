package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/wallet/base"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"os"
	"path/filepath"

	//"strings"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"log"
)

func GenSeed() []string {
	const (
		grpcHost = "127.0.0.1:10009"
	)
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	// Load the TLS certificate
	cert, err := ioutil.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}

	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建 WalletUnlocker 客户端
	client := lnrpc.NewWalletUnlockerClient(conn)

	passphrase := ""
	var aezeedPassphrase = []byte(passphrase)

	// Generating random seed entropy:
	seedEntropy := make([]byte, 16)
	_, err = rand.Read(seedEntropy)
	if err != nil {
		log.Fatalf("could not generate seed entropy: %v", err)
	}

	// 准备 WalletUnlocker 请求
	request := &lnrpc.GenSeedRequest{
		AezeedPassphrase: aezeedPassphrase,
		SeedEntropy:      seedEntropy,
	}

	// 调用 InitWallet gRPC 方法
	response, err := client.GenSeed(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling InitWallet: %v", err)
	}

	// 处理 gRPC 响应
	return response.CipherSeedMnemonic
}

func Initwallet(Psw string) bool {
	const (
		grpcHost = "127.0.0.1:10009"
	)
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	// Load the TLS certificate
	cert, err := ioutil.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}

	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建 WalletUnlocker 客户端
	client := lnrpc.NewWalletUnlockerClient(conn)

	passphrase := ""
	var aezeedPassphrase = []byte(passphrase)

	request := &lnrpc.InitWalletRequest{
		WalletPassword:     []byte(Psw),
		CipherSeedMnemonic: GenSeed(),
		AezeedPassphrase:   aezeedPassphrase,
	}

	response, err := client.InitWallet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling InitWallet: %v", err)
	}

	d1 := response.AdminMacaroon

	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	err = os.MkdirAll(newFilePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error calling InitWallet: %v", err)
	}
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")

	f, err := os.Create(macaroonPath)
	if err != nil {
		log.Fatalf("Error calling InitWallet: %v", err)
		return false
	}
	_, err = f.Write(d1)
	if err != nil {
		f.Close()
		log.Fatalf("Error calling InitWallet: %v", err)

		return false
	}
	log.Println("successful")
	err = f.Close()
	if err != nil {
		log.Fatalf("Error calling InitWallet: %v", err)

		return false
	}

	return true
}

func Unlockwallet(Psw string) bool {
	const (
		grpcHost = "127.0.0.1:10009"
	)

	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	// Load the TLS certificate
	cert, err := ioutil.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}

	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建 WalletUnlocker 客户端
	client := lnrpc.NewWalletUnlockerClient(conn)

	request := &lnrpc.UnlockWalletRequest{
		WalletPassword: []byte(Psw),
	}
	_, err = client.UnlockWallet(context.Background(), request)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
	_currentPasswordByteSlice, _ := hex.DecodeString(currentPassword)
	_newPasswordByteSlice, _ := hex.DecodeString(newPassword)
	request := &lnrpc.ChangePasswordRequest{
		CurrentPassword: _currentPasswordByteSlice,
		NewPassword:     _newPasswordByteSlice,
	}
	response, err := client.ChangePassword(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ChangePassword err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}
