package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetBlockByMempoolByMempool() {}

func GetBlockHeaderByMempool() {}

func GetBlockHeightByMempool() {}

func GetBlockTimestampByMempool() {}

func GetBlockRawByMempool() {}

func GetBlockStatusByMempool() {}

func GetBlockTipHeightByMempool() string {
	targetUrl := "https://mempool.space/testnet/api/blocks/tip/height"
	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("%s http.Get :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "http get fail.", "")
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var height string
	height = string(bodyBytes)
	return MakeJsonResult(true, "", height)
}

// BlockTipHeight
// @dev: NOT STANDARD RESULT RETURN
func BlockTipHeight() int {
	targetUrl := "https://mempool.space/testnet/api/blocks/tip/height"
	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("%s http.Get :%v\n", GetTimeNow(), err)
		return 0
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	height, _ := strconv.Atoi(string(bodyBytes))
	return height
}

func GetBlockTipHashByMempool() {}

func GetBlockTransactionIDByMempool() {}

func GetBlockTransactionIDsByMempool() {}

func GetBlockTransactionsByMempool() {}

func GetBlocksByMempool() {}

func GetBlocksBulkByMempool() {}
