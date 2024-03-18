package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetStateForSubscribe() bool {
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
	client := lnrpc.NewStateClient(conn)
	request := &lnrpc.SubscribeStateRequest{}
	response, err := client.SubscribeState(context.Background(), request)
	log.Print(response)
	return true
}
