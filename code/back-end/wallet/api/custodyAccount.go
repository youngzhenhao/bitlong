package api

import (
	"encoding/json"
	"fmt"
	"github.com/wallet/api/apipost"
)

func ApplyInvoiceRequest(amount int64, memo string, token string) (string, error) {
	user := "testuser"
	pass := "testpass"

	token, err := apipost.Login(user, pass)
	if err != nil {
		return "", err
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
	request, err := apipost.SendPostRequest(url, token, requestBody)
	if err != nil {
		return "", err
	}
	invoice := struct {
		Error   string `json:"error"`
		Invoice string `json:"invoice"`
	}{}
	err = json.Unmarshal(request, &invoice)
	if err != nil {
		return "", err
	}
	if invoice.Error != "" {
		return "", fmt.Errorf(invoice.Error)
	}
	return invoice.Invoice, nil
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
	url := apipost.Server + apipost.Pay
	apipost.SendPostRequest(url, token, requestBody)
}
