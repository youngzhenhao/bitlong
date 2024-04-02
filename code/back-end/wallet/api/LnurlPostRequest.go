package api

import (
	"encoding/json"
	"fmt"
	"github.com/wallet/base"
	"io"
	"net/http"
	"net/url"
)

type InvoiceResponse struct {
	Time    string `json:"time"`
	ID      string `json:"id"`
	Amount  string `json:"amount"`
	Invoice string `json:"invoice"`
	Result  bool   `json:"result"`
}

type UserResponse struct {
	Time       string `json:"time"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Socket     string `json:"socket"`
	RemotePort string `json:"remote_port"`
	Result     bool   `json:"result"`
	Lnurl      string `json:"lnurl"`
}

func PostServerToUploadUserInfo(id, name, localPort, remotePort string) string {

	serverDomainOrSocket := base.QueryConfigByKey("LnurlServerHost")
	targetUrl := "http://" + serverDomainOrSocket + "/upload/user"

	payload := url.Values{"id": {id}, "name": {name}, "local_port": {localPort}, "remote_port": {remotePort}}

	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)

	var userResponse UserResponse
	if err := json.Unmarshal(bodyBytes, &userResponse); err != nil {
		fmt.Printf("%s PSTUUI json.Unmarshal :%v\n", GetTimeNow(), err)
		return ""
	}
	return userResponse.Lnurl
}

// PostPhoneToAddInvoice called by server
func PostPhoneToAddInvoice(remotePort, amount string) string {

	frpsForwardSocket := fmt.Sprintf("%s:%s", base.QueryConfigByKey("serverAddr"), remotePort)

	targetUrl := "http://" + frpsForwardSocket + "/addInvoice"

	payload := url.Values{"amount": {amount}}

	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)

	var invoiceResponse InvoiceResponse
	if err := json.Unmarshal(bodyBytes, &invoiceResponse); err != nil {
		fmt.Printf("%s PPTAI json.Unmarshal :%v\n", GetTimeNow(), err)
		return ""
	}
	return invoiceResponse.Invoice
}

// PostServerToPayByPhoneAddInvoice called by Bob
func PostServerToPayByPhoneAddInvoice(lnu, amount string) string {
	targetUrl := Decode(lnu)
	payload := url.Values{"amount": {amount}}
	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var invoiceResponse InvoiceResponse
	if err := json.Unmarshal(bodyBytes, &invoiceResponse); err != nil {
		fmt.Printf("%s PSTPBPAI json.Unmarshal :%v\n", GetTimeNow(), err)
		return ""
	}
	return invoiceResponse.Invoice
}
