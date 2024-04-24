package lnd

import (
	"fmt"
	"github.com/wallet/api"
)

const (
	WalletEmpty    string = "[NoWallet]"
	WalletUnlocked string = "[Unlocked]"
	WalletLocked   string = "[Unlocked]"
)

func LndWorkflowCreateWallet(password string) bool {
	fmt.Printf("%s %sStarting to generate random mnemonics to create wallet...\n", api.GetTimeNow(), WalletEmpty)
	seed := GenSeed()
	fmt.Printf("%s Cipher seed mnemonic have been generated:\n", api.GetTimeNow())
	fmt.Printf("%s Initializing wallet with this mnemonic and password...\n", api.GetTimeNow())
	result := InitWallet(seed, password)
	if result {
		fmt.Printf("%s %sWallet created successfully!\n", api.GetTimeNow(), WalletUnlocked)
	}
	fmt.Printf("%s %sWallet created!\n", api.GetTimeNow(), WalletUnlocked)
	return result
}

func LndWorkflowUnlockWallet(password string) bool {
	fmt.Printf("%s %sUnlocking...\n", api.GetTimeNow(), WalletLocked)
	result := UnlockWallet(password)
	if result {
		fmt.Printf("%s %sWallet unlocked successfully!\n", api.GetTimeNow(), WalletUnlocked)
	}
	return result
}

func LndWorkflowChangeWalletPassword(currentPassword string, newPassword string) bool {
	fmt.Printf("%s %sChanging password...\n", api.GetTimeNow(), WalletLocked)
	result := ChangePassword(currentPassword, newPassword)
	if result {
		fmt.Printf("%s %sPassword changed successfully!\n", api.GetTimeNow(), WalletUnlocked)
	}
	fmt.Printf("%s %sPassword changed!\n", api.GetTimeNow(), WalletUnlocked)
	return result
}

func LndWorkflowGetNewAddressAndWalletBalance() {

}
