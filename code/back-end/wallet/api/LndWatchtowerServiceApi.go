package api

import (
	"context"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc/watchtowerrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"path/filepath"
)

// WatchtowerGetInfo
//
//	@Description: GetInfo returns general information concerning the companion watchtower including its public key and URIs where the server is currently listening for clients.
//	@return string
func WatchtowerGetInfo() string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
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
	client := watchtowerrpc.NewWatchtowerClient(conn)
	request := &watchtowerrpc.GetInfoRequest{}
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		fmt.Printf("%s watchtowerrpc GetInfo err: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}
