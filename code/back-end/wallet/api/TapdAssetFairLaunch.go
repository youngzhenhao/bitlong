package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IssuanceHistoryInfo struct {
	AssetName    string `json:"asset_name"`
	AssetID      string `json:"asset_id"`
	AssetType    int    `json:"asset_type"`
	IssuanceTime int    `json:"issuance_time"`
	State        int    `json:"state"`
}

func GetIssuanceTransactionByteSize() int {
	// TODO: need to complete
	return 170
}

func GetMintTransactionByteSize() int {
	// TODO: need to complete
	return 170
}

// TODO: Assemble local and server asset issuance data
// @dev: Use new makeJsonResult

// http://127.0.0.1:8080/v1/fair_launch/query/own_set

func GetOwnSet(token string) (string, error) {
	url := serverHost + "/v1/fair_launch/query/own_set"
	// 创建HTTP请求
	responce, err := SendGetReq(url, token, nil)
	return responce, err
}

func GetRate(token string) (string, error) {
	url := serverHost + "/v1/fee/query/rate"
	responce, err := SendGetReq(url, token, nil)
	return responce, err

}

func GetAssetQueryMint(token string, FairLaunchInfoId string, MintedNumber int) (string, error) {
	url := serverHost + "/v1/fee/query/rate"
	resquest := struct {
		FairLaunchInfoId string `json:"fair_launch_info_id"`
		MintedNumber     int    `json:"minted_number"`
	}{}
	requestBody, _ := json.Marshal(resquest)
	// 创建HTTP请求
	responce, err := SendGetReq(url, token, requestBody)
	return responce, err
}
func SendGetReq(url string, token string, requestBody []byte) (string, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("An error occurred while creating an HTTP request:", err)
	}
	// 设置Authorization Header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("An error occurred while sending an HTTP request:", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("An error occurred while closing the HTTP response body:", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
