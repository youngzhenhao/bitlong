package api

import (
	"fmt"
)

const (
	WalletEmpty    string = "[NoWallet]"
	WalletUnlocked string = "[Unlocked]"
	WalletLocked   string = "[Unlocked]"
)

func LndWorkflowCreateWallet(password string) bool {
	fmt.Printf("%s %sStarting to generate random mnemonics to create wallet...\n", GetTimeNow(), WalletEmpty)
	seed := GenSeed()
	fmt.Printf("%s Cipher seed mnemonic have been generated:\n", GetTimeNow())
	fmt.Printf("%s Initializing wallet with this mnemonic and password...\n", GetTimeNow())
	result := InitWallet(seed, password)
	if result {
		fmt.Printf("%s %sWallet created successfully!\n", GetTimeNow(), WalletUnlocked)
	}
	fmt.Printf("%s %sWallet created!\n", GetTimeNow(), WalletUnlocked)
	return result
}

func LndWorkflowUnlockWallet(password string) bool {
	fmt.Printf("%s %sUnlocking...\n", GetTimeNow(), WalletLocked)
	result := UnlockWallet(password)
	if result {
		fmt.Printf("%s %sWallet unlocked successfully!\n", GetTimeNow(), WalletUnlocked)
	}
	return result
}

func LndWorkflowChangeWalletPassword(currentPassword string, newPassword string) bool {
	fmt.Printf("%s %sChanging password...\n", GetTimeNow(), WalletLocked)
	result := ChangePassword(currentPassword, newPassword)
	if result {
		fmt.Printf("%s %sPassword changed successfully!\n", GetTimeNow(), WalletUnlocked)
	}
	fmt.Printf("%s %sPassword changed!\n", GetTimeNow(), WalletUnlocked)
	return result
}

func LndWorkflowGetNewAddressAndWalletBalance() {

}
