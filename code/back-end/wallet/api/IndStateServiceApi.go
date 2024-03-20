package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
)

func GetStateForSubscribe() bool {
	grpcHost := base.QueryConfigByKey("lndhost")
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
	client := lnrpc.NewStateClient(conn)
	request := &lnrpc.SubscribeStateRequest{}
	response, err := client.SubscribeState(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc NewAddress err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// GetState
//
//	@Description: 返回当前钱包状态，不做进一步更改
//	@return *lnrpc.GetStateResponse
//
// func GetState() *lnrpc.GetStateResponse {
func GetState() string {
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
	client := lnrpc.NewStateClient(conn)
	request := &lnrpc.GetStateRequest{}
	response, err := client.GetState(context.Background(), request)
	if err != nil {
		log.Printf("watchtowerrpc GetState err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}
