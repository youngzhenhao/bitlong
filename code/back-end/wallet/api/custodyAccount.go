package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func sendPostRequest(url string, token string, requestBody []byte) {

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建HTTP请求时出错:", err)
		return
	}

	// 设置Authorization Header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送POST请求时出错:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭HTTP响应体时出错:", err)
		}
	}(resp.Body)
}

func ApplyInvoiceRequest(amount int64, memo string) {

	//apply := struct {
	//	Amount int64  `json:"amount"`
	//	Memo   string `json:"memo"`
	//}{}

}
