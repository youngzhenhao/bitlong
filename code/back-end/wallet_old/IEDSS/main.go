package main

//
//import (
//	"fmt"
//	"github.com/btcsuite/btcd/chaincfg"
//	"github.com/ethereum/go-ethereum/accounts"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/tyler-smith/go-bip39"
//)
//
//func main() {
//	// Assuming the mnemonic is already generated and is passed here securely
//	mnemonic := "your existing mnemonic phrase here" // Replace with your actual mnemonic
//
//	// Generate a seed from the mnemonic
//	seed := bip39.NewSeed(mnemonic, "")
//
//	// Define the derivation path (e.g., standard BIP44 path for Bitcoin on Mainnet)
//	// m/44'/0'/0'/0 for the first address in the main account
//	path := accounts.DefaultBaseDerivationPath
//
//	// Create a master key from the seed
//	masterKey, err := hd.NewMaster(seed, &chaincfg.MainNetParams)
//	if err != nil {
//		fmt.Println("Error generating master key:", err)
//		return
//	}
//
//	// Derive the first child key using the specified path
//	childKey, err := masterKey.Derive(path)
//	if err != nil {
//		fmt.Println("Error deriving the child key:", err)
//		return
//	}
//
//	// Convert the child key into an ECDSA private key
//	privateKeyECDSA, err := crypto.ToECDSA(childKey.Key)
//	if err != nil {
//		fmt.Println("Error converting to ECDSA:", err)
//		return
//	}
//
//	// Extract the public key
//	publicKey := privateKeyECDSA.Public()
//	publicKeyECDSA, ok := publicKey.(*crypto.PublicKey)
//	if !ok {
//		fmt.Println("Error asserting type:", err)
//		return
//	}
//
//	// Convert the public key to a hexadecimal string
//	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
//	publicKeyHex := fmt.Sprintf("%x", publicKeyBytes)
//
//	fmt.Println("Derived Public Key:", publicKeyHex)
//}
