package litd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/wallet/api"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
)

func SubServerStatus() string {
	response, err := subServerStatus()
	if err != nil {
		return api.MakeJsonResult(false, err.Error(), nil)
	}
	return api.MakeJsonResult(true, "", response)
}

func GetTapdStatus() bool {
	response, err := subServerStatus()
	if err != nil {
		return false
	}
	if len(response.SubServers) == 0 {
		return false
	}
	return response.SubServers["taproot-assets"].Running
}

func GetLitStatus() bool {
	response, err := subServerStatus()
	if err != nil {
		return false
	}
	if len(response.SubServers) == 0 {
		return false
	}
	return response.SubServers["lit"].Running
}

func subServerStatus() (*litrpc.SubServerStatusResp, error) {
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
		fmt.Printf("%s Failed to read cert file: %s", api.GetTimeNow(), err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf(api.GetTimeNow() + "Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(api.NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", api.GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", api.GetTimeNow(), err)
		}
	}(conn)

	client := litrpc.NewStatusClient(conn)
	request := &litrpc.SubServerStatusReq{}
	response, err := client.SubServerStatus(context.Background(), request)
	return response, err
}
