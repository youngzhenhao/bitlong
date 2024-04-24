package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetAddressResponse struct {
	Address    string `json:"address"`
	ChainStats struct {
		FundedTxoCount int   `json:"funded_txo_count"`
		FundedTxoSum   int64 `json:"funded_txo_sum"`
		SpentTxoCount  int   `json:"spent_txo_count"`
		SpentTxoSum    int   `json:"spent_txo_sum"`
		TxCount        int   `json:"tx_count"`
	} `json:"chain_stats"`
	MempoolStats struct {
		FundedTxoCount int `json:"funded_txo_count"`
		FundedTxoSum   int `json:"funded_txo_sum"`
		SpentTxoCount  int `json:"spent_txo_count"`
		SpentTxoSum    int `json:"spent_txo_sum"`
		TxCount        int `json:"tx_count"`
	} `json:"mempool_stats"`
}

func GetAddressInfoByMempool(address string) string {

	targetUrl := "https://mempool.space/testnet/api/address/" + address
	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "http get fail.", "")
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var getAddressResponse GetAddressResponse
	if err := json.Unmarshal(bodyBytes, &getAddressResponse); err != nil {
		fmt.Printf("%s PSTUUI json.Unmarshal :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "Unmarshal response body fail.", "")
	}
	return MakeJsonResult(true, "", getAddressResponse)
}

func GetAddressTransactions() {}

func GetAddressTransactionsChain() {}

func GetAddressTransactionsMempool() {}

func GetAddressUTXO() {}

func GetAddressValidation() {}
