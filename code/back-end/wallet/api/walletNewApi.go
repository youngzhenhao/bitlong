package api

import (
	"encoding/json"
	"github.com/wallet/address"
	"github.com/wallet/config"
)

// 返回数据结构
type AddrData struct {
	LegacyAddress      string `json:"legacyAddress"`      //老地址
	LegacyPriKey       string `json:"legacyPriKey"`       //老地址私钥
	NativeSegWitAddr   string `json:"nativeSegWitAddr"`   //原生地址
	NativeSegWitPriKey string `json:"nativeSegWitPriKey"` //原生地址私钥
	NestedSegWitAddr   string `json:"nestedSegWitAddr"`   //原生兼容地址
	NestedSegWitPriKey string `json:"nestedSegWitPriKey"` //原生兼容地址私钥
}

func GetAddr(Mnemonic string, Account int64, addressIndex int64) string {
	getK := &address.GetKeys{
		CoinType:     int64(config.BTC),
		Mnemonic:     Mnemonic,
		Account:      Account,
		AddressIndex: addressIndex,
	}
	master, _ := address.HdNewKey(getK) //构造注入参数
	LegacyAddress, _ := master.LegacyAddress()
	LegacyPriKey, _ := master.GetPriKey(1)
	NativeSegWitAddr, _ := master.NativeSegWitAddr()
	NativeSegWitPriKey, _ := master.GetPriKey(2)
	NestedSegWitAddr, _ := master.NestedSegWitAddr()
	NestedSegWitPriKey, _ := master.GetPriKey(3)
	data := &AddrData{
		LegacyAddress:      LegacyAddress,
		LegacyPriKey:       LegacyPriKey,
		NativeSegWitAddr:   NativeSegWitAddr,
		NativeSegWitPriKey: NativeSegWitPriKey,
		NestedSegWitAddr:   NestedSegWitAddr,
		NestedSegWitPriKey: NestedSegWitPriKey,
	}
	dataStr, _ := json.Marshal(data)
	return RegReturn(200, 10003, string(dataStr), "GetAddr successful")
}
