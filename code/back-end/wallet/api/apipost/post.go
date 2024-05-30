package apipost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// 申请发票请求地址
	Server     = "http://localhost:8081"
	Apply      = "/custodyAccount/invoice/apply"
	Pay        = "/custodyAccount/invoice/pay"
	LoginUrl   = "/login"
	RefreshUrl = "/refresh"
)

func SendPostRequest(url string, token string, requestBody []byte) ([]byte, error) {

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("An error occurred while creating an HTTP request:", err)
		return nil, err
	}

	// 设置Authorization Header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("An error occurred while sending a POST request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("An error occurred while closing the HTTP response body:", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

type User struct {
	Username string `gorm:"unique;column:user_name" json:"userName"` // 正确地将unique和column选项放在同一个gorm标签内
	Password string `gorm:"column:password" json:"password"`
}

func Login(username string, password string) (string, error) {
	user := User{
		Username: username,
		Password: password,
	}
	requestBody, _ := json.Marshal(user)
	a, err := SendPostRequest(Server+LoginUrl, "", requestBody)
	if err != nil {
		return "", err
	}
	result := struct {
		Error string `json:"error"`
		Token string `json:"token"`
	}{}
	err = json.Unmarshal(a, &result)
	if err != nil {
		fmt.Println("An error occurred while unmarshalling the response body:", err)
	}
	if result.Error != "" {
		return "", fmt.Errorf(result.Error)
	}
	return result.Token, err
}

func Refresh(username string, password string) (string, error) {
	user := User{
		Username: username,
		Password: password,
	}
	requestBody, _ := json.Marshal(user)
	a, err := SendPostRequest(Server+RefreshUrl, "", requestBody)
	if err != nil {
		return "", err
	}
	result := struct {
		Error string `json:"error"`
		Token string `json:"token"`
	}{}
	err = json.Unmarshal(a, &result)
	if err != nil {
		fmt.Println("An error occurred while unmarshalling the response body:", err)
	}
	if result.Error != "" {
		return "", fmt.Errorf(result.Error)
	}
	return result.Token, err
}
