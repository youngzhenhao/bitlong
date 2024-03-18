package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/wallet/other"
	"github.com/wallet/query"
	"github.com/wallet/tx"
)

// 返回数据结构
type RegMsgDataRes struct {
	MsgId int
	Code  int
	Data  string
	Msg   string
}
type Fee struct {
	Slow    float64 `json:"slow"`    // 慢速手续费
	Base    float64 `json:"base"`    // 正常手续费
	Fast    float64 `json:"fast"`    // 快速手续费
	FeeUnit string  `json:"feeUnit"` // 手续费单位
}

// GetBalance 获取余额
func GetBalance(address string) string {
	balance, err := other.GetBalance("bitcoin", "mainnet", address)
	if err != nil {
		return RegReturn(400, 10002, "0", "GetBalance error")
	}
	return RegReturn(200, 10002, other.CoinStringDiv(balance, 8), "GetBalance successful")
}

func GetFee() string {
	gas, err := other.GetGasFee("bitcoin", "mainnet")
	if err != nil {
		return RegReturn(400, 10003, "", "GetGasFee error")
	}
	data := &Fee{
		Fast:    float64(gas.EstimatedFees.Medium * 320),
		Base:    float64(gas.EstimatedFees.Medium * 200),
		Slow:    float64(gas.EstimatedFees.Medium * 120),
		FeeUnit: "sal",
	}
	dataStr, err := json.Marshal(data)
	fmt.Println(err)
	return RegReturn(200, 10003, string(dataStr), "GetFee successful")
}

// TransactionList 交易列表
func TransactionList(address string, size int) string {
	list, err := other.GetTXList("bitcoin", "mainnet", address, "", size)
	if err != nil {
		return RegReturn(400, 10004, "", "GetTXList error")
	}
	ret := query.TxListReqData(list, address)
	dataStr, _ := json.Marshal(ret)
	return RegReturn(200, 10004, string(dataStr), "GetTxList successful")
}

// Transfer 转账
func Transfer(fromAddress, toAddress, priKey string, amount int64, fee int64) string {
	// 获取utxo
	utxos, err := other.GetUnspentUTXO(fromAddress, amount+fee, "bitcoin")
	if err != nil {
		return RegReturn(400, 10005, "", "get utxo error")
	}
	isWitness := true
	if string(fromAddress[0]) != "3" {
		isWitness = false
	} else if string(fromAddress[0]) == "b" {
		return RegReturn(400, 10005, "", "unknown address type")
	}
	signTx, err := tx.CreateTx(priKey, toAddress, amount, fee, utxos, isWitness) //btc创建交易
	if err != nil {
		return RegReturn(400, 10005, "", "CreateTx  error")
	}
	tx, err := other.SendTx("bitcoin", "mainnet", signTx)
	if err != nil {
		return RegReturn(400, 10005, "", "SendTx  error")
	}
	if len(tx.TxId) == 0 {
		return RegReturn(400, 10005, "", "SendTx  error")
	}
	return RegReturn(200, 10005, tx.TxId, "Broadcast successful")

}

func RegReturn(Code int, MsgId int, Data string, msg string) string {
	frameMsgDataRes := &RegMsgDataRes{
		Code:  Code,
		MsgId: MsgId,
		Data:  Data,
		Msg:   msg,
	}
	regData, _ := json.Marshal(frameMsgDataRes)
	return string(regData)
}

func GenerateMnemonic() string {
	bytes := 256 / 8
	// 生成熵源
	entropy := make([]byte, bytes)
	_, err := rand.Read(entropy)
	if err != nil {
		return RegReturn(400, 10007, "", "GenerateMnemonic error")
	}

	// 使用熵源生成助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return RegReturn(400, 10007, "", "NewMnemonic error")
	}
	return RegReturn(200, 10007, mnemonic, "GenerateMnemonic successful")
}
