package api

import (
	"context"
	"fmt"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/wallet/api/connect"
)

func SubServerStatus() string {
	response, err := subServerStatus()
	if err != nil {
		return MakeJsonErrorResult(DefaultErr, err.Error(), nil)
	}
	return MakeJsonErrorResult(SUCCESS, "", response)
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
	conn, clearUp, err := connect.GetConnection("litd", true)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()

	client := litrpc.NewStatusClient(conn)
	request := &litrpc.SubServerStatusReq{}
	response, err := client.SubServerStatus(context.Background(), request)
	return response, err
}
