package service

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/nbd-wtf/go-nostr"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"log"
)

const (
	keyId = "keyInfoId"
)

type PkInfo struct {
	Pubkey  string `json:"pubkey"`
	NpubKey string `json:"npubKey"`
}

func GenerateKeys(mnemonic, passphrase string) (string, error) {
	retrievedKey, err := readDb()
	if err != nil {
		fmt.Println("err:", err)
	}
	if retrievedKey != nil {
		publicKeyHex := fmt.Sprintf("%064x", retrievedKey.PublicKey)
		fmt.Println("pub:", publicKeyHex)
		return publicKeyHex, nil
	}
	fmt.Println("mnemonic:", mnemonic)
	seed := bip39.NewSeed(mnemonic, passphrase)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", err
	}
	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	if err != nil {
		return "", err
	}
	privKey, _ := btcec.PrivKeyFromBytes(childKey.Key)

	privateKeyHex := fmt.Sprintf("%064x", privKey.Serialize())
	publicKeyHex, err := getPublicKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	fmt.Println("pub1:", publicKeyHex)
	err1 := saveDb(privateKeyHex, publicKeyHex)
	if err1 != nil {
		fmt.Println("err:", err1)
		return "", fmt.Errorf("failed to save keys: %s", err1)
	} // Assuming the bucket name is "Keys"
	return publicKeyHex, nil
}

func saveDb(privateKeyHex string, publicKeyHex string) error {
	db, err := InitDB() // 确保已经写了这个函数来初始化数据库和表
	if err != nil {
		fmt.Printf("Failed to initialize the database: %s\n", err)
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	keyStore := KeyStore{DB: db}

	keyInfo := KeyInfo{
		ID:         keyId, // This should be dynamically set or managed
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}
	key, err := keyStore.ReadKey(keyId)
	if err != nil {
		fmt.Printf("Failed to read data: %s\n", err)
	}
	if key == nil {
		err := keyStore.CreateOrUpdateKey(&keyInfo)
		if err != nil {
			fmt.Printf("Failed to save data: %s\n", err)
			return err
		}
	}
	return nil
}
func sign(privateKeyHex string, message string) (string, error) {
	var evt nostr.Event
	err := json.Unmarshal([]byte(message), &evt)
	if err != nil {
		return "", err
	}
	if err := evt.Sign(privateKeyHex); err != nil {
		return "", err
	}
	marshal, err := json.Marshal(evt)
	if err != nil {
		return "", err
	}
	return string(marshal), nil

}
func readDb() (*KeyInfo, error) {
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %s\n", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	keyStore := &KeyStore{DB: db}
	// 调用 ReadKey 获取特定的密钥
	keyInfo, err := keyStore.ReadKey(keyId)
	if err != nil {
		log.Printf("Failed to read key %s: %s", keyId, err)
		return nil, err
	} else {
		fmt.Printf("Key: %+v\n", keyInfo)
		return keyInfo, nil
	}
}

func SignMessage(message string) (string, error) {
	retrievedKey, err := readDb()
	if err != nil {
		fmt.Printf("err is :%v\n", err)
	}
	signInfo, err := sign(retrievedKey.PrivateKey, message)
	if err != nil {
		return "", err
	}

	fmt.Println(signInfo)
	return signInfo, nil
}
func getPublicKey(sk string) (string, error) {
	b, err := hex.DecodeString(sk)
	if err != nil {
		return "", err
	}
	_, pk := btcec.PrivKeyFromBytes(b)
	return hex.EncodeToString(schnorr.SerializePubKey(pk)), nil
}

func getNoStrAddress(pk string) (string, error) {
	compressedPubKeyBytes, err := hex.DecodeString(pk)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return "", err
	}
	// 将公钥编码为 Base58
	base58EncodedPubKey := base58.Encode(compressedPubKeyBytes)
	// 添加 nostr 协议所需的前缀
	nostrPubKey := "npub" + base58EncodedPubKey
	fmt.Println("Nostr address:", nostrPubKey)
	return nostrPubKey, nil
}

func GetPublicKey() (string, string, error) {
	retrievedKey, err := readDb()
	if err != nil {
		fmt.Printf("err is :%v\n", err)
		return "", "", err
	}
	publicKeyHex := fmt.Sprintf("%064X", retrievedKey.PublicKey)
	fmt.Println("publicKeyHex", publicKeyHex)
	address, err := getNoStrAddress(publicKeyHex)
	if err != nil {
		return "", "", err
	}
	return publicKeyHex, address, nil
}

func GetJsonPublicKey() (string, error) {
	var pkInfo PkInfo
	retrievedKey, err := readDb()
	if err != nil {
		fmt.Printf("err is :%v\n", err)
		return "", err
	}
	publicKeyHex := fmt.Sprintf("%064X", retrievedKey.PublicKey)
	fmt.Println("publicKeyHex", publicKeyHex)
	address, err := getNoStrAddress(publicKeyHex)
	if err != nil {
		return "", err
	}
	pkInfo.Pubkey = publicKeyHex
	pkInfo.NpubKey = address
	marshal, err := json.Marshal(pkInfo)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
