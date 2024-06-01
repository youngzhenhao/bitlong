package api

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/api/connect"
	"github.com/wallet/models"
	"io"
	"net/http"
)

type IssuanceHistoryInfo struct {
	IsFairLaunchIssuance bool   `json:"isFairLaunchIssuance"`
	AssetName            string `json:"asset_name"`
	AssetID              string `json:"asset_id"`
	AssetType            int    `json:"asset_type"`
	IssuanceTime         int    `json:"issuance_time"`
	State                int    `json:"state"`
}

func GetIssuanceTransactionFee(token string) (fee int, err error) {
	size := GetIssuanceTransactionByteSize()
	serverFeeRateResponse, err := GetServerFeeRate(token)
	if err != nil {
		LogError("", err)
		return 0, err
	}
	feeRate := serverFeeRateResponse.Data.SatPerB
	return feeRate * size, err
}

func GetMintTransactionByteFee(token string, id int, number int) (fee int, err error) {
	size := GetMintTransactionByteSize()
	serverQueryMintResponse, err := GetServerQueryMint(token, id, number)
	if err != nil {
		LogError("", err)
		return 0, err
	}
	feeRate := serverQueryMintResponse.Data.CalculatedFeeRateSatPerB
	return feeRate * size, err
}

func GetIssuanceTransactionByteSize() int {
	// TODO: need to complete
	return 170
}

func GetMintTransactionByteSize() int {
	// TODO: need to complete
	return 170
}

type ServerOwnSetFairLaunchInfoResponse struct {
	Success bool                    `json:"success"`
	Error   string                  `json:"error"`
	Data    []models.FairLaunchInfo `json:"data"`
}

func GetServerOwnSetFairLaunchInfos(token string) (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	serverDomainOrSocket := "132.232.109.84:8090"
	//serverDomainOrSocket := "127.0.0.1:8080"
	url := "http://" + serverDomainOrSocket + "/v1/fair_launch/query/own_set"
	client := &http.Client{}
	var jsonData []byte
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			LogError("", err)
		}
	}(response.Body)
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	var ownSetFairLaunchInfos ServerOwnSetFairLaunchInfoResponse
	if err := json.Unmarshal(bodyBytes, &ownSetFairLaunchInfos); err != nil {
		fmt.Printf("%s json.Unmarshal :%v\n", GetTimeNow(), err)
		return nil, err
	}
	return &ownSetFairLaunchInfos.Data, nil
}

// TODO: Assemble local and server asset issuance data
// @dev: Use new makeJsonResult

// http://127.0.0.1:8080/v1/fair_launch/query/own_set

func ProcessOwnSetFairLaunchResponseToIssuanceHistoryInfo(fairLaunchInfos *[]models.FairLaunchInfo) (*[]IssuanceHistoryInfo, error) {
	var err error
	var issuanceHistoryInfos []IssuanceHistoryInfo
	if fairLaunchInfos == nil || len(*fairLaunchInfos) == 0 {
		err = errors.New("fairLaunchInfos is null")
		LogError("", err)
		return nil, err
	}
	for _, fairLaunchInfo := range *fairLaunchInfos {
		issuanceHistoryInfos = append(issuanceHistoryInfos, IssuanceHistoryInfo{
			IsFairLaunchIssuance: true,
			AssetName:            fairLaunchInfo.Name,
			AssetID:              fairLaunchInfo.AssetID,
			AssetType:            int(fairLaunchInfo.AssetType),
			IssuanceTime:         fairLaunchInfo.IssuanceTime,
			State:                int(fairLaunchInfo.State),
		})
	}
	return &issuanceHistoryInfos, nil

}

type ServerFeeRateResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    struct {
		SatPerKw int     `json:"sat_per_kw"`
		SatPerB  int     `json:"sat_per_b"`
		BtcPerKb float64 `json:"btc_per_kb"`
	} `json:"data"`
}

func GetServerFeeRate(token string) (serverFeeRateResponse *ServerFeeRateResponse, err error) {
	serverDomainOrSocket := "132.232.109.84:8090"
	//serverDomainOrSocket := "127.0.0.1:8080"
	url := "http://" + serverDomainOrSocket + "/v1/fee/query/rate"
	client := &http.Client{}
	var jsonData []byte
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			LogError("", err)
		}
	}(response.Body)
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	if err = json.Unmarshal(bodyBytes, &serverFeeRateResponse); err != nil {
		fmt.Printf("%s json.Unmarshal :%v\n", GetTimeNow(), err)
		return nil, err
	}
	return serverFeeRateResponse, nil
}

type ServerQueryMintResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    struct {
		CalculatedFeeRateSatPerB  int  `json:"calculated_fee_rate_sat_per_b"`
		CalculatedFeeRateSatPerKw int  `json:"calculated_fee_rate_sat_per_kw"`
		InventoryAmount           int  `json:"inventory_amount"`
		IsMintAvailable           bool `json:"is_mint_available"`
	} `json:"data"`
}

func GetServerQueryMint(token string, id int, number int) (serverQueryMintResponse *ServerQueryMintResponse, err error) {
	serverDomainOrSocket := "132.232.109.84:8090"
	//serverDomainOrSocket := "127.0.0.1:8080"
	url := "http://" + serverDomainOrSocket + "/v1/fair_launch/query/mint"
	client := &http.Client{}
	requestJson := struct {
		FairLaunchInfoId int `json:"fair_launch_info_id"`
		MintedNumber     int `json:"minted_number"`
	}{
		FairLaunchInfoId: id,
		MintedNumber:     number,
	}
	requestJsonBytes, _ := json.Marshal(requestJson)
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(requestJsonBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			LogError("", err)
		}
	}(response.Body)
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	if err = json.Unmarshal(bodyBytes, &serverQueryMintResponse); err != nil {
		fmt.Printf("%s json.Unmarshal :%v\n", GetTimeNow(), err)
		return nil, err
	}
	return serverQueryMintResponse, nil
}

// GetServerIssuanceHistoryInfos
// @Description: Get Server Issuance History Info
// @param token
// @return *[]IssuanceHistoryInfo
// @return error
func GetServerIssuanceHistoryInfos(token string) (*[]IssuanceHistoryInfo, error) {
	fairLaunchInfos, err := GetServerOwnSetFairLaunchInfos(token)
	if err != nil {
		LogError("", err)
		return nil, err
	}
	issuanceHistoryInfos, err := ProcessOwnSetFairLaunchResponseToIssuanceHistoryInfo(fairLaunchInfos)
	if err != nil {
		LogError("", err)
		return nil, err
	}
	return issuanceHistoryInfos, nil
}

func GetLocalIssuanceHistoryInfos() {
	//	TODO: need to complete
}

func ListBatchesAndGetResponse() (*mintrpc.ListBatchResponse, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := mintrpc.NewMintClient(conn)
	request := &mintrpc.ListBatchRequest{}
	response, err := client.ListBatches(context.Background(), request)
	if err != nil {
		fmt.Printf("%s mintrpc ListBatches Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, nil
}

func GetTransactionsAndGetResponse() (*lnrpc.TransactionDetails, error) {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.GetTransactionsRequest{}
	response, err := client.GetTransactions(context.Background(), request)
	return response, err
}

func ListAssetAndGetResponse() (*[]ListAssetResponse, error) {
	response, err := listAssets(false, true, false)
	if err != nil {
		LogError("", err)
		return nil, err
	}
	var listAssetResponses []ListAssetResponse
	for _, asset := range (*response).Assets {
		listAssetResponses = append(listAssetResponses, ListAssetResponse{
			AssetGenesis: AssetGenesisStruct{
				GenesisPoint: asset.AssetGenesis.GenesisPoint,
				Name:         asset.AssetGenesis.Name,
				MetaHash:     hex.EncodeToString(asset.AssetGenesis.MetaHash),
				AssetID:      hex.EncodeToString(asset.AssetGenesis.AssetId),
			},
			Amount:           int64(asset.Amount),
			ScriptKey:        hex.EncodeToString(asset.ScriptKey),
			ScriptKeyIsLocal: asset.ScriptKeyIsLocal,
			ChainAnchor: ChainAnchorStruct{
				AnchorTx:        hex.EncodeToString(asset.ChainAnchor.AnchorTx),
				AnchorBlockHash: asset.ChainAnchor.AnchorBlockHash,
				AnchorOutpoint:  asset.ChainAnchor.AnchorOutpoint,
				InternalKey:     hex.EncodeToString(asset.ChainAnchor.InternalKey),
				MerkleRoot:      hex.EncodeToString(asset.ChainAnchor.MerkleRoot),
			}})
	}
	return &listAssetResponses, nil
}

type AssetGenesisStruct struct {
	GenesisPoint string `json:"genesis_point"`
	Name         string `json:"name"`
	MetaHash     string `json:"meta_hash"`
	AssetID      string `json:"asset_id"`
}

type ChainAnchorStruct struct {
	AnchorTx        string `json:"anchor_tx"`
	AnchorBlockHash string `json:"anchor_block_hash"`
	AnchorOutpoint  string `json:"anchor_outpoint"`
	InternalKey     string `json:"internal_key"`
	MerkleRoot      string `json:"merkle_root"`
}

type ListAssetResponse struct {
	AssetGenesis     AssetGenesisStruct `json:"asset_genesis"`
	Amount           int64              `json:"amount"`
	ScriptKey        string             `json:"script_key"`
	ScriptKeyIsLocal bool               `json:"script_key_is_local"`
	ChainAnchor      ChainAnchorStruct  `json:"chain_anchor"`
}

func GetTimestampByBatchTxidWithGetTransactionsResponse(transactionDetails *lnrpc.TransactionDetails, batchTxid string) (timestamp int, err error) {
	for _, transaction := range transactionDetails.Transactions {
		if batchTxid == transaction.TxHash {
			return int(transaction.TimeStamp), nil
		}
	}
	err = errors.New("transaction not found")
	return 0, err
}

func GetAssetIdByBatchTxidWithListAssetResponse(listAssetResponse *taprpc.ListAssetResponse, batchTxid string) (assetId string, err error) {
	for _, asset := range listAssetResponse.Assets {
		tx, _ := getTransactionAndIndexByOutpoint(asset.ChainAnchor.AnchorOutpoint)
		if batchTxid == tx {
			return hex.EncodeToString(asset.AssetGenesis.AssetId), nil
		}
	}
	err = errors.New("asset not found")
	return "", err
}

func ProcessListAssetAndGetResponseToIssuanceHistoryInfo(listAssetResponse *[]ListAssetResponse) (*[]IssuanceHistoryInfo, error) {
	var err error
	var issuanceHistoryInfos []IssuanceHistoryInfo
	batchs, err := ListBatchesAndGetResponse()
	if err != nil {
		LogError("", err)
		return nil, err
	}
	transactions, err := GetTransactionsAndGetResponse()
	if err != nil {
		LogError("", err)
		return nil, err
	}
	assets, err := ListAssetAndGetResponse()
	if err != nil {
		LogError("", err)
		return nil, err
	}
	var timestamp int
	var assetId string
	for _, batch := range (*batchs).Batches {
		timestamp, err = GetTimestampByBatchTxidWithGetTransactionsResponse(transactions, batch.BatchTxid)
		if err != nil {
			LogError("", err)
			//	@dev:do not return
		}
		assetId, err = GetAssetIdByBatchTxidWithListAssetResponse(assets, batch.BatchTxid)
	}

}
