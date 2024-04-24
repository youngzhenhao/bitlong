package api

import (
	"encoding/json"
	"fmt"
	"github.com/wallet/base"
	"io"
	"net/http"
	"net/url"
)

type AvailablePortResponse struct {
	Time       string `json:"time"`
	RemotePort uint16 `json:"remote_port"`
	Result     bool   `json:"result"`
}

type IsPortListeningResponse struct {
	Time       string `json:"time"`
	RemotePort string `json:"remote_port"`
	Result     bool   `json:"result"`
	Listening  bool   `json:"listening"`
}

func RequestServerGetPortAvailable(socket string) int {

	serverDomainOrSocket := socket
	targetUrl := "http://" + serverDomainOrSocket + "/availablePort"

	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("%s http Get err :%v\n", GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var availablePortResponse AvailablePortResponse
	if err := json.Unmarshal(bodyBytes, &availablePortResponse); err != nil {
		fmt.Printf("%s RSGPA json.Unmarshal :%v\n", GetTimeNow(), err)
		return 0
	}
	return int(availablePortResponse.RemotePort)
}

func RequestPostServerIsPortListening(remotePort string) bool {
	serverDomainOrSocket := base.QueryConfigByKey("LnurlServerHost")
	targetUrl := "http://" + serverDomainOrSocket + "/isPortListening"
	payload := url.Values{"remote_port": {remotePort}}
	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var isPortListeningResponse IsPortListeningResponse
	if err := json.Unmarshal(bodyBytes, &isPortListeningResponse); err != nil {
		fmt.Printf("%s RPSIPL json.Unmarshal :%v\n", GetTimeNow(), err)
		return true
	}
	return isPortListeningResponse.Listening
}
