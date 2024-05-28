package rpcclient

import (
	"context"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/wallet/api/connect"
	"time"
)

func getTaprootAssetsClient() (taprpc.TaprootAssetsClient, func(), error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	client := taprpc.NewTaprootAssetsClient(conn)
	return client, clearUp, nil
}

func DecodeAddr(addr string) (*taprpc.Addr, error) {

	client, clearUp, err := getTaprootAssetsClient()
	if err != nil {
		return nil, err
	}
	defer clearUp()

	request := &taprpc.DecodeAddrRequest{
		Addr: addr,
	}
	response, err := client.DecodeAddr(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc DecodeAddr Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, nil
}

func GetTimeNow() any {
	return time.Now().Format("2006/01/02 15:04:05")
}
