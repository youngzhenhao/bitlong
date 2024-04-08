package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
)

func subServerStatus() *litrpc.SubServerStatusResp {
	grpcHost := base.QueryConfigByKey("litdhost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lit"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "lit.macaroon")
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
		fmt.Printf(GetTimeNow() + "Failed to append cert")
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

	client := litrpc.NewStatusClient(conn)
	request := &litrpc.SubServerStatusReq{}
	response, err := client.SubServerStatus(context.Background(), request)
	if err != nil {
		fmt.Printf("%s subserver status err: %v\n", GetTimeNow(), err)
	}
	return response
}
func GetTapdStatus() bool {
	litstatus := subServerStatus().GetSubServers()
	if len(litstatus) == 0 {
		return false
	}
	return litstatus["taproot-assets"].Running
}

func GetLitStatus() bool {
	litstatus := subServerStatus().GetSubServers()
	if len(litstatus) == 0 {
		return false
	}
	return litstatus["lit"].Running
}
