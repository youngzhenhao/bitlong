package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wallet/models"
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
			AssetName:    fairLaunchInfo.Name,
			AssetID:      fairLaunchInfo.AssetID,
			AssetType:    int(fairLaunchInfo.AssetType),
			IssuanceTime: fairLaunchInfo.IssuanceTime,
			State:        int(fairLaunchInfo.State),
		})
	}
	return &issuanceHistoryInfos, nil

}

func GetOwnSet(token string) (string, error) {
	url := serverHost + "/v1/fair_launch/query/own_set"
	// Create an HTTP request
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
	// Create an HTTP request
	responce, err := SendGetReq(url, token, requestBody)
	return responce, err
}
func SendGetReq(url string, token string, requestBody []byte) (string, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("An error occurred while creating an HTTP request:", err)
	}
	// Set Authorization Header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	// Send HTTP request
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
