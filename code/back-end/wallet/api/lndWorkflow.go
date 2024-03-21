package api

import "log"

type WalletStatus string

const (
	WALLET_EMPTY    WalletStatus = "[NoWallet]"
	WALLET_UNLOCKED WalletStatus = "[Unlocked]"
	WALLET_LOCKED   WalletStatus = "[Unlocked]"
)

func LndWorkflowCreateWallet(password string) bool {
	log.Println(WALLET_EMPTY, "Starting to generate random mnemonics to create wallet...")
	seed := GenSeed()
	log.Println("Cipher seed mnemonic have been generated:\n", seed)
	log.Println("Initializing wallet with this mnemonic and password...")
	result := InitWallet(seed, password)
	if result {
		log.Println(WALLET_UNLOCKED, "Wallet created successfully!")
	}
	log.Println(WALLET_UNLOCKED, "Wallet created!")
	return result
}

func LndWorkflowUnlockWallet(password string) bool {
	log.Println(WALLET_LOCKED, "Unlocking...")
	result := UnlockWallet(password)
	if result {
		log.Println(WALLET_UNLOCKED, "Wallet unlocked successfully!")
	}
	return result
}

func LndWorkflowChangeWalletPassword(currentPassword string, newPassword string) bool {
	log.Println(WALLET_LOCKED, "Changing password...")
	result := ChangePassword(currentPassword, newPassword)
	if result {
		log.Println(WALLET_UNLOCKED, "Password changed successfully!")
	}
	log.Println(WALLET_UNLOCKED, "Password changed!")
	return result
}

func LndWorkflowGetNewAddressAndWalletBalance() {

}
