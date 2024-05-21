package api

import (
	"context"
	"fmt"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"path/filepath"
)

func SubServerStatus() string {
	response, err := subServerStatus()
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
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
	newFilePath := filepath.Join(base.Configure("lit"), base.NetWork)
	macaroonPath := filepath.Join(newFilePath, "lit.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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

	client := litrpc.NewStatusClient(conn)
	request := &litrpc.SubServerStatusReq{}
	response, err := client.SubServerStatus(context.Background(), request)
	return response, err
}
