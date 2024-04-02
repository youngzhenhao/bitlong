package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AvailablePortResponse struct {
	Time       string `json:"time"`
	RemotePort uint16 `json:"remote_port"`
	Result     bool   `json:"result"`
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
