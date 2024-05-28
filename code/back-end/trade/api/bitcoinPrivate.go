package api

import (
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
	"trade/config"
)

func getBitcoinConnConfig() *rpcclient.ConnConfig {
	return &rpcclient.ConnConfig{
		Host:         fmt.Sprintf("%s:%d", config.GetLoadConfig().ApiConfig.Bitcoin.Host, config.GetLoadConfig().ApiConfig.Bitcoin.Port),
		User:         config.GetLoadConfig().ApiConfig.Bitcoin.RpcUser,
		Pass:         config.GetLoadConfig().ApiConfig.Bitcoin.RpcPasswd,
		HTTPPostMode: config.GetLoadConfig().ApiConfig.Bitcoin.HTTPPostMode,
		DisableTLS:   config.GetLoadConfig().ApiConfig.Bitcoin.DisableTLS,
	}
}

func estimateSmartFee(confTarget int64, mode *btcjson.EstimateSmartFeeMode) (feeResult *btcjson.EstimateSmartFeeResult, err error) {
	connCfg := getBitcoinConnConfig()
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		fmt.Println("Error connecting to the RPC server:", err)
		return
	}
	defer client.Shutdown()
	feeResult, err = client.EstimateSmartFee(confTarget, mode)
	if err != nil {
		fmt.Println("Error calling EstimateSmartFeeAndGetResult:", err)
		return
	}
	return feeResult, nil
}
