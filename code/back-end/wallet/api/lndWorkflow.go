package api

import (
	"fmt"
)

type WalletStatus string

const (
	WALLET_EMPTY    WalletStatus = "[NoWallet]"
	WALLET_UNLOCKED WalletStatus = "[Unlocked]"
	WALLET_LOCKED   WalletStatus = "[Unlocked]"
)

func LndWorkflowCreateWallet(password string) bool {
	fmt.Printf("%s %sStarting to generate random mnemonics to create wallet...\n", GetTimeNow(), WALLET_EMPTY)
	seed := GenSeed()
	fmt.Printf("%s Cipher seed mnemonic have been generated:\n", GetTimeNow())
	fmt.Printf("%s Initializing wallet with this mnemonic and password...\n", GetTimeNow())
	result := InitWallet(seed, password)
	if result {
		fmt.Printf("%s %sWallet created successfully!\n", GetTimeNow(), WALLET_UNLOCKED)
	}
	fmt.Printf("%s %sWallet created!\n", GetTimeNow(), WALLET_UNLOCKED)
	return result
}

func LndWorkflowUnlockWallet(password string) bool {
	fmt.Printf("%s %sUnlocking...\n", GetTimeNow(), WALLET_LOCKED)
	result := UnlockWallet(password)
	if result {
		fmt.Printf("%s %sWallet unlocked successfully!\n", GetTimeNow(), WALLET_UNLOCKED)
	}
	return result
}

func LndWorkflowChangeWalletPassword(currentPassword string, newPassword string) bool {
	fmt.Printf("%s %sChanging password...\n", GetTimeNow(), WALLET_LOCKED)
	result := ChangePassword(currentPassword, newPassword)
	if result {
		fmt.Printf("%s %sPassword changed successfully!\n", GetTimeNow(), WALLET_UNLOCKED)
	}
	fmt.Printf("%s %sPassword changed!\n", GetTimeNow(), WALLET_UNLOCKED)
	return result
}

func LndWorkflowGetNewAddressAndWalletBalance() {

}
