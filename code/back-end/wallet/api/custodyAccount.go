package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// 申请发票请求地址
	Server = "http://localhost:8080"
	Apply  = "/api/v1/custody/apply-invoice"
	Pay    = "/api/v1/custody/pay-invoice"
)

func sendPostRequest(url string, token string, requestBody []byte) {

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("An error occurred while creating an HTTP request:", err)
		return
	}

	// 设置Authorization Header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("An error occurred while sending a POST request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("An error occurred while closing the HTTP response body:", err)
		}
	}(resp.Body)
}

func ApplyInvoiceRequest(amount int64, memo string, token string) {
	applyRequest := struct {
		Amount int64  `json:"amount"`
		Memo   string `json:"memo"`
	}{
		Amount: amount,
		Memo:   memo,
	}
	requestBody, _ := json.Marshal(applyRequest)
	url := Server + Apply
	sendPostRequest(url, token, requestBody)
}

func PayInvoiceRequest(invoiceId string, FeeLimit int64, token string) {
	payRequest := struct {
		Invoice  string `json:"invoice"`
		FeeLimit int64  `json:"feeLimit"`
	}{
		Invoice:  invoiceId,
		FeeLimit: FeeLimit,
	}
	requestBody, _ := json.Marshal(payRequest)
	url := Server + Pay
	sendPostRequest(url, token, requestBody)
}
