package api

import (
	"context"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/api/connect"
)

// GetStateForSubscribe
//
//	@Description:SubscribeState subscribes to the state of the wallet.
//	The current wallet state will always be delivered immediately.
//	@return bool
func GetStateForSubscribe() bool {
	conn, clearUp, err := connect.GetConnection("lnd", true)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := lnrpc.NewStateClient(conn)
	request := &lnrpc.SubscribeStateRequest{}
	response, err := client.SubscribeState(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc SubscribeState err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

func GetState() string {
	response, err := getState()
	if err != nil {
		fmt.Printf("%s watchtowerrpc GetState err: %v\n", GetTimeNow(), err)
		return "NO_START_LND"
	}
	return response.State.String()
}

// getState
//
//	@Description: GetState returns the current wallet state without streaming further changes.
//	@return string
func getState() (*lnrpc.GetStateResponse, error) {
	conn, clearUp, err := connect.GetConnection("lnd", true)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := lnrpc.NewStateClient(conn)
	request := &lnrpc.GetStateRequest{}
	response, err := client.GetState(context.Background(), request)
	return response, err
}
