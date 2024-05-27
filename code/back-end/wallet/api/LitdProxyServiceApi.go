package api

import (
	"context"
	"fmt"
	terminal "github.com/lightninglabs/lightning-terminal"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/wallet/api/connect"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"path/filepath"
)

func LitdStopDaemon() bool {
	_, err := litdStopDaemon()
	if err != nil {
		fmt.Printf("%s litrpc StopRequest err: %v\n", GetTimeNow(), err)
		return false
	}
	return true
}

func litdStopDaemon() (*litrpc.StopDaemonResponse, error) {
	grpcHost := base.QueryConfigByKey("litdhost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lit"), base.NetWork)
	macaroonPath := filepath.Join(newFilePath, "lit.macaroon")
	creds := connect.NewTlsCert(tlsCertPath)
	macaroon := connect.GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(connect.NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := litrpc.NewProxyClient(conn)
	request := &litrpc.StopDaemonRequest{}
	response, err := client.StopDaemon(context.Background(), request)
	return response, err
}

func LitdLocalStop() {
	fmt.Printf("%s User stop Litd at local...\n", GetTimeNow())
	terminal.StopLitd()
}
