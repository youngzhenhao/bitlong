package api

import (
	"encoding/json"
	"fmt"
)

const (
	// 申请发票请求地址
	Apply           = "/custodyAccount/invoice/apply"
	Pay             = "/custodyAccount/invoice/pay"
	QuerybalanceUrl = "/custodyAccount/invoice/querybalance"
)

func ApplyInvoiceRequest(amount int64, memo string, token string) ([]byte, error) {
	applyRequest := struct {
		Amount int64  `json:"amount"`
		Memo   string `json:"memo"`
	}{
		Amount: amount,
		Memo:   memo,
	}
	requestBody, _ := json.Marshal(applyRequest)
	url := Server + Apply
	response, err := SendPostRequest(url, token, requestBody)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func PayInvoiceRequest(invoiceId string, FeeLimit int64, token string) ([]byte, error) {
	payRequest := struct {
		Invoice  string `json:"invoice"`
		FeeLimit int64  `json:"feeLimit"`
	}{
		Invoice:  invoiceId,
		FeeLimit: FeeLimit,
	}
	requestBody, _ := json.Marshal(payRequest)
	url := Server + Pay
	response, err := SendPostRequest(url, token, requestBody)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response, nil
}

// TODO:
func QueryInvoicesRequest(invoiceId string, token string) ([]byte, error) {

	return nil, nil
}

// TODO:
func QueryPaymentsRequest(token string) ([]byte, error) {
	return nil, nil
}

func QueryBalanceRequest(token string) ([]byte, error) {
	url := Server + QuerybalanceUrl
	response, err := SendPostRequest(url, token, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response, nil
}
