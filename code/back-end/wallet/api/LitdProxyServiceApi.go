package api

import (
	"context"
	"fmt"
	terminal "github.com/lightninglabs/lightning-terminal"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/wallet/api/connect"
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
	conn, clearUp, err := connect.GetConnection("litd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()

	client := litrpc.NewProxyClient(conn)
	request := &litrpc.StopDaemonRequest{}
	response, err := client.StopDaemon(context.Background(), request)
	return response, err
}

func LitdLocalStop() {
	fmt.Printf("%s User stop Litd at local...\n", GetTimeNow())
	terminal.StopLitd()
}
