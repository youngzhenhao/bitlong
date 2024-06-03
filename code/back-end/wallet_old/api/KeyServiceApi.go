package api

import (
	"fmt"
	"github.com/wallet/service"
)

func GenerateKeys(mnemonic string) string {
	keys, err := service.GenerateKeys(mnemonic, "")
	if err != nil {
		fmt.Printf("GenerateKeys->errl1:%x", err)
		return ""
	}
	publicKeyHex := fmt.Sprintf("%x", keys)
	fmt.Printf("publ1:%x", publicKeyHex)
	fmt.Printf("errl1:%x", err)
	return publicKeyHex
}

func GetPublicKey() string {
	pb, _, err := service.GetPublicKey()
	if err != nil {
		fmt.Printf("GetPublicKey->errl1:%x", err)
		return ""
	}
	return pb
}
func GetNPublicKey() string {
	_, nPub, err := service.GetPublicKey()
	if err != nil {
		fmt.Printf("GetNPublicKey->errl1:%x", err)
		return ""
	}
	return nPub
}
func GetJsonPublicKey() string {
	keyInfo, err := service.GetJsonPublicKey()
	if err != nil {
		fmt.Printf("GetNPublicKey->errl1:%x", err)
		return ""
	}
	return keyInfo
}

func SignMess(message string) string {
	sign, err := service.SignMessage(message)
	if err != nil {
		fmt.Printf("SignMess->errl1:%x", err)
		return ""
	}
	return sign
}
