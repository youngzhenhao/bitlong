package other

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// httpGet get请求
func httpGetParam(url string, params map[string]string) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 添加请求头
	//req.Header.Add("x-apiKey", config.ApiKey)
	//req.Header.Add("Ok-Access-Key", config.OkAccessKey)
	query := req.URL.Query()
	if params != nil && len(params) != 0 {
		for k, v := range params {
			if v != "" {
				query.Add(k, v)
			}
		}
	}
	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	//fmt.Println("请求返回：", string(body), resp.StatusCode)
	if resp.StatusCode != 200 || err != nil {
		//time.Sleep(300 * time.Millisecond)
		return nil, err
	}
	return body, nil
	//	ch <- 1
}

// HttpGet get请求
func HttpGet(url string) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		//engine.Log.Info("GetClient NewRequest error: %s", err.Error())
		return nil, nil
	}
	client.Timeout = 6 * time.Second
	// 添加请求头
	resp, err := client.Do(req)
	if err != nil {
		//engine.Log.Info("GetClient Do error: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 || err != nil {
		return nil, err
	}
	return body, nil
	//	ch <- 1
}

// httpPost post请求
func HttpPost(url string, data map[string]interface{}) ([]byte, error) {
	bytesData, _ := json.Marshal(data)
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 || err != nil {
		return nil, errors.New(string(respBytes))
	}
	//byte数组直接转成string，优化内存
	return respBytes, err
}
