package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc/credentials"
	"io/ioutil"

	//"strings"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func Test_main2(t *testing.T) {
	const (
		grpcHost    = "127.0.0.1:10009"
		tlsCertPath = "tls.cert"
	)
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
	log.Println(response)
}
