package api

import (
	"encoding/json"
	"github.com/wallet/address"
	"github.com/wallet/config"
)

type AddrData struct {
	LegacyAddress      string `json:"legacyAddress"`
	LegacyPriKey       string `json:"legacyPriKey"`
	NativeSegWitAddr   string `json:"nativeSegWitAddr"`
	NativeSegWitPriKey string `json:"nativeSegWitPriKey"`
	NestedSegWitAddr   string `json:"nestedSegWitAddr"`
	NestedSegWitPriKey string `json:"nestedSegWitPriKey"`
}

func GetAddr(Mnemonic string, Account int64, addressIndex int64) string {
	getK := &address.GetKeys{
		CoinType:     int64(config.BTC),
		Mnemonic:     Mnemonic,
		Account:      Account,
		AddressIndex: addressIndex,
	}
	master, _ := address.HdNewKey(getK)
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
