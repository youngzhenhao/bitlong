package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightningnetwork/lnd/lnrpc/chainrpc"
	"github.com/wallet/api/connect"
)

func GetBlockWrap(blockHash string) string {
	response, err := GetBlock(blockHash)
	if err != nil {
		return MakeJsonErrorResult(DefaultErr, err.Error(), nil)
	}
	msgBlock := &wire.MsgBlock{}
	blockReader := bytes.NewReader(response.RawBlock)
	err = msgBlock.Deserialize(blockReader)
	return MakeJsonErrorResult(SUCCESS, "", msgBlock)
}

func GetBlock(blockHashStr string) (response *chainrpc.GetBlockResponse, err error) {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()

	blockHash, err := chainhash.NewHashFromStr(blockHashStr)
	client := chainrpc.NewChainKitClient(conn)
	request := &chainrpc.GetBlockRequest{
		BlockHash: blockHash.CloneBytes(),
	}
	response, err = client.GetBlock(context.Background(), request)
	return response, err
}
