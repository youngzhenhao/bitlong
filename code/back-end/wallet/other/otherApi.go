package other

import (
	"encoding/json"
	"fmt"
	"github.com/wallet/config"
)

func GetBalance(chain string, netWork string, address string) (string, error) {
	var regRet []BalanceRegData
	retData, err := getApiData(chain, netWork, address, config.BlockDamonGetBalanceUrl)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal([]byte(retData), &regRet)
	if err != nil {
		return "", err
	}
	return regRet[0].ConfirmedBalance, nil
}

func getApiData(chain string, netWork string, address string, getUrl string) (string, error) {
	url := fmt.Sprintf(getUrl, chain, netWork, address)
	retData, err := HttpGet(config.BlockDamonApiUrl + url)
	if err != nil {
		return "", err
	}
	return string(retData), nil
}

// GetTXList 获取交易列表
func GetTXList(chain string, netWork string, address string, pageToken string, size int) (*TxListData, error) {
	var regRet TxListData
	retData, err := getApiDataForTxList(chain, netWork, address, config.BlockDamonTxListUrl, pageToken, size)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(retData), &regRet)
	if err != nil {
		return nil, err
	}
	return &regRet, nil
}

func getApiDataForTxList(chain string, netWork string, address string, getUrl string, pageToken string, size int) (string, error) {
	url := fmt.Sprintf(getUrl, chain, netWork, address, pageToken, size)
	retData, err := HttpGet(config.BlockDamonApiUrl + url)
	if err != nil {
		return "", err
	}
	return string(retData), nil
}

func GetGasFee(chain string, netWork string) (*GasFeeData, error) {
	var regRet GasFeeData
	url := fmt.Sprintf(config.BlockDamonGetGasFeeUrl, chain, netWork)
	retData, err := HttpGet(config.BlockDamonApiUrl + url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(retData, &regRet)
	if err != nil {
		return nil, err
	}
	return &regRet, nil
}

// SendTx 广播交易
func SendTx(chain string, netWork string, txHash string) (*SendTxRegData, error) {
	var regRet *SendTxRegData
	//接收请求参数
	paramsData := make(map[string]interface{})
	paramsData["tx"] = txHash
	//发起请求
	url := fmt.Sprintf(config.BlockDamonTxSendUrl, chain, netWork)
	retData, err := HttpPost(config.BlockDamonApiUrl+url, paramsData)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(retData, &regRet)
	if err != nil {
		return nil, err
	}
	return regRet, nil
}

// GetUTXO 获取UTXO btc ltc bch
func GetUTXO(chain string, netWork string, address string) ([]UTXOdata, error) {
	var regRet UTXORegData
	url := fmt.Sprintf(config.BlockDamonGetUTXOUrl, chain, netWork, address)
	retData, err := HttpGet(config.BlockDamonApiUrl + url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(retData, &regRet)
	if err != nil {
		return nil, err
	}
	return regRet.Data, nil
}
