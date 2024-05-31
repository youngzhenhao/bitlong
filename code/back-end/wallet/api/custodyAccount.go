package api

import (
	"encoding/json"
	"github.com/wallet/api/apipost"
)

func ApplyInvoiceRequest(amount int64, memo string, token string) ([]byte, error) {
	username := "testuser"
	password := "testpass"
	token, err := apipost.Login(username, password)
	if err != nil {
		return nil, err
	}

	applyRequest := struct {
		Amount int64  `json:"amount"`
		Memo   string `json:"memo"`
	}{
		Amount: amount,
		Memo:   memo,
	}
	requestBody, _ := json.Marshal(applyRequest)
	url := apipost.Server + apipost.Apply
	response, err := apipost.SendPostRequest(url, token, requestBody)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func PayInvoiceRequest(invoiceId string, FeeLimit int64, token string) ([]byte, error) {
	username := "testuser"
	password := "testpass"
	token, err := apipost.Login(username, password)
	if err != nil {
		return nil, err
	}

	payRequest := struct {
		Invoice  string `json:"invoice"`
		FeeLimit int64  `json:"feeLimit"`
	}{
		Invoice:  invoiceId,
		FeeLimit: FeeLimit,
	}
	requestBody, _ := json.Marshal(payRequest)
	url := apipost.Server + apipost.Pay
	response, err := apipost.SendPostRequest(url, token, requestBody)
	return response, err
}

func QueryInvoicesRequest(invoiceId string, token string) ([]byte, error) {
	return nil, nil
}

func QueryPaymentsRequest(token string) ([]byte, error) {
	return nil, nil
}

func QueryBalanceRequest(token string) ([]byte, error) {
	return nil, nil
}
