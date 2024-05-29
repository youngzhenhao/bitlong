package rpcclient

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/wallet/api/connect"
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

func AddrReceives() (*taprpc.AddrReceivesResponse, error) {
	client, clearUp, err := getTaprootAssetsClient()
	if err != nil {
		return nil, err
	}
	defer clearUp()

	request := &taprpc.AddrReceivesRequest{}
	response, err := client.AddrReceives(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
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

func QueryAddr() (*taprpc.QueryAddrResponse, error) {
	client, clearUp, err := getTaprootAssetsClient()
	if err != nil {
		return nil, err
	}
	defer clearUp()
	request := &taprpc.QueryAddrRequest{}
	response, err := client.QueryAddrs(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc QueryAddr Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, nil
}

func NewAddr(assetId string, amt int) (*taprpc.Addr, error) {
	client, clearUp, err := getTaprootAssetsClient()
	if err != nil {
		return nil, err
	}
	defer clearUp()

	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	request := &taprpc.NewAddrRequest{
		AssetId: _assetIdByteSlice,
		Amt:     uint64(amt),
	}
	response, err := client.NewAddr(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func ListTransfers() (*taprpc.ListTransfersResponse, error) {
	client, clearUp, err := getTaprootAssetsClient()
	if err != nil {
		return nil, err
	}
	defer clearUp()

	request := &taprpc.ListTransfersRequest{}
	response, err := client.ListTransfers(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, err
}
