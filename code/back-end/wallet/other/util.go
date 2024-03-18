package other

import (
	"errors"
	"github.com/wallet/tx"
	"math"
	"strconv"
	"strings"
)

func CoinStringDiv(val string, power int) string {
	mi := power + 1
	if len(val) < mi {
		val = strings.Repeat("0", mi-len(val)) + val
	}
	return val[0:len(val)-power] + "." + val[len(val)-power:]
}

func ParseEvent(data []EventData, regData *TransactionBase, address string, coinName string) {
	fromValue := float64(0)
	toValue := float64(0)
	for _, value := range data {
		switch value.Type {
		case "transfer":
			regData.From = value.From
			regData.To = value.To
			regData.Amount = CoinStringDiv(strconv.FormatFloat(value.Amount, 'f', -1, 64), 6)
			if coinName == "sol" {
				regData.Amount = CoinStringDiv(strconv.FormatFloat(value.Amount, 'f', -1, 64), 9)
			}
		case "fee":
			regData.Fee = CoinStringDiv(strconv.FormatFloat(value.Amount, 'f', -1, 64), 8)
			switch coinName {
			case "sol":
				regData.Fee = CoinStringDiv(strconv.FormatFloat(value.Amount, 'f', -1, 64), 9)
				break
			case "xrp":
				regData.Fee = CoinStringDiv(strconv.FormatFloat(value.Amount, 'f', -1, 64), 6)
				break
			}
		case "utxo_input":
			regData.UtxoInput = append(regData.UtxoInput, &UtxoAmount{
				Address: value.From,
				Amount:  strconv.FormatFloat(value.Amount, 'f', -1, 64),
			})
			if strings.EqualFold(address, value.From) {
				fromValue += value.Amount
			}
		case "utxo_output":
			regData.UtxoOutput = append(regData.UtxoOutput, &UtxoAmount{
				Address: value.To,
				Amount:  strconv.FormatFloat(value.Amount, 'f', -1, 64),
			})

			if strings.EqualFold(address, value.To) {
				toValue += value.Amount
			}
		}
	}
	if regData.Amount == "" {
		num, _ := strconv.ParseFloat(regData.Fee, 64)
		fees := num * 1e8
		//btc bch ltc doge除1/e8
		regData.Amount = CoinStringDiv(strconv.FormatFloat(math.Abs(fromValue-toValue-fees), 'f', -1, 64), 8)
	}
	if strings.EqualFold(address, regData.From) || fromValue > toValue {
		regData.Send = true
	}
}

// GetUnspentUTXO 获取指定数量的未花费的UTXO
func GetUnspentUTXO(address string, value int64, coinType string) ([]*tx.Unspent, error) {
	var (
		bal     int64
		ok      bool
		unspent = make([]*tx.Unspent, 0, 1)
	)
	// get balance
	utxoRet, err := GetUTXO(coinType, "mainnet", address)
	if err != nil {
		return nil, err
	}
	for _, v := range utxoRet {
		if !v.IsSpent { //未发送
			bal += v.Value
			unspent = append(unspent, &tx.Unspent{
				TxHash: v.Mined.TxId,
				Index:  uint32(v.Mined.Index),
				Value:  v.Value,
			})
			if v.Value > value || bal >= value {
				ok = true
				break
			}
		}
	}
	if !ok {
		return nil, errors.New("lack of balance")
	}
	return unspent, nil
}
