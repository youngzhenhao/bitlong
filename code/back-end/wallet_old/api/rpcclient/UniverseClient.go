package rpcclient

import (
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"github.com/wallet/api/connect"
)

func getUniverseClient() (*universerpc.UniverseClient, func(), error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	client := universerpc.NewUniverseClient(conn)
	return &client, clearUp, nil
}
