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

type GetAddressTransactionsResponse []struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid    string `json:"txid"`
		Vout    int    `json:"vout"`
		Prevout struct {
			Scriptpubkey        string `json:"scriptpubkey"`
			ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
			ScriptpubkeyType    string `json:"scriptpubkey_type"`
			ScriptpubkeyAddress string `json:"scriptpubkey_address"`
			Value               int    `json:"value"`
		} `json:"prevout"`
		Scriptsig    string   `json:"scriptsig"`
		ScriptsigAsm string   `json:"scriptsig_asm"`
		Witness      []string `json:"witness"`
		IsCoinbase   bool     `json:"is_coinbase"`
		Sequence     int64    `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Scriptpubkey        string `json:"scriptpubkey"`
		ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
		ScriptpubkeyType    string `json:"scriptpubkey_type"`
		ScriptpubkeyAddress string `json:"scriptpubkey_address"`
		Value               int    `json:"value"`
	} `json:"vout"`
	Size   int `json:"size"`
	Weight int `json:"weight"`
	Sigops int `json:"sigops"`
	Fee    int `json:"fee"`
	Status struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight int    `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   int    `json:"block_time"`
	} `json:"status"`
}

// GetAddressInfoByMempool
// @Description: Get address info by mempool api
// @param address
// @return string
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
		fmt.Printf("%s GAIBM json.Unmarshal :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "Unmarshal response body fail.", "")
	}
	return MakeJsonResult(true, "", getAddressResponse)
}

// GetAddressTransactionsByMempool
// @Description: Get address transactions by mempool api
// @param address
// @return string
func GetAddressTransactionsByMempool(address string) string {
	targetUrl := "https://mempool.space/testnet/api/address/" + address + "/txs"
	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("%s http.Get :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "http get fail.", "")
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var getAddressTransactionsResponse GetAddressTransactionsResponse
	if err := json.Unmarshal(bodyBytes, &getAddressTransactionsResponse); err != nil {
		fmt.Printf("%s GATBM json.Unmarshal :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "Unmarshal response body fail.", "")
	}
	return MakeJsonResult(true, "", getAddressTransactionsResponse)
}

func GetAddressTransactionsChainByMempool() {}

func GetAddressTransactionsMempoolByMempool() {}

func GetAddressUTXOByMempool() {}

func GetAddressValidationByMempool() {}
