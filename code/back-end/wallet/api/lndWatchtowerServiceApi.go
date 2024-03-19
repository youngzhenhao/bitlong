package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"

	"github.com/lightningnetwork/lnd/lnrpc/watchtowerrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
)

// WatchtowerGetInfo
//
//	@Description: 返回有关同伴监视塔的一般信息，包括其公钥和服务器当前正在监听客户端的 URI
//	@return *watchtowerrpc.GetInfoResponse
//
// func WatchtowerGetInfo() *watchtowerrpc.GetInfoResponse {
func WatchtowerGetInfo() string {
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
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("conn Close err: %v", err)
		}
	}(conn)
	client := watchtowerrpc.NewWatchtowerClient(conn)
	request := &watchtowerrpc.GetInfoRequest{}
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		log.Printf("watchtowerrpc GetInfo err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}
