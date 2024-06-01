package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Deprecated
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
