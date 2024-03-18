package address

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/cosmos/go-bip39"
	"github.com/tyler-smith/go-bip32"
	"github.com/wallet/config"
	"sync"
)

type KeyManager struct {
	mnemonic   string
	passphrase string
	keys       map[string]*bip32.Key
	mux        sync.Mutex
}
type BipKey struct {
	path     string
	bip32Key *bip32.Key
}

func NewKeyManager(mnemonic string) (*KeyManager, error) {
	km := &KeyManager{
		mnemonic:   mnemonic,
		passphrase: "",
		keys:       make(map[string]*bip32.Key, 0),
	}
	return km, nil
}

// LegacyAddress 老地址
func (k *Key) LegacyAddress() (string, error) {
	km, _ := NewKeyManager(k.Opt.Mnemonic)
	key, _ := km.GetKey(config.PurposeBIP44, config.BTC, uint32(k.Opt.Account), 0, uint32(k.Opt.AddressIndex))
	prvKey, _ := btcec.PrivKeyFromBytes(key.bip32Key.Key)
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return "", err
	}
	//生成正常p2pkh地址
	serializedPubKey := btcwif.SerializePubKey()
	//从pubkey哈希生成正常的p2wkh地址
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressPubKeyHash(witnessProg, k.Opt.Params)
	if err != nil {
		return "", err
	}
	address := addressWitnessPubKeyHash.EncodeAddress()
	return address, nil
}

// NativeSegWitAddr 原生隔离见证地址
func (k *Key) NativeSegWitAddr() (string, error) {
	km, _ := NewKeyManager(k.Opt.Mnemonic)
	key, _ := km.GetKey(config.PurposeBIP84, config.BTC, uint32(k.Opt.Account), 0, uint32(k.Opt.AddressIndex))
	prvKey, _ := btcec.PrivKeyFromBytes(key.bip32Key.Key)
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return "", err
	}
	//生成正常p2pkh地址
	serializedPubKey := btcwif.SerializePubKey()
	//从pubkey哈希生成正常的p2wkh地址
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, k.Opt.Params)
	if err != nil {
		return "", err
	}
	nestedSegWitAddr := addressWitnessPubKeyHash.EncodeAddress()
	return nestedSegWitAddr, nil
}

// NestedSegWitAddr 隔离见证兼容
func (k *Key) NestedSegWitAddr() (string, error) {
	km, _ := NewKeyManager(k.Opt.Mnemonic)
	key, _ := km.GetKey(config.PurposeBIP49, config.BTC, uint32(k.Opt.Account), 0, uint32(k.Opt.AddressIndex))
	prvKey, _ := btcec.PrivKeyFromBytes(key.bip32Key.Key)
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return "", err
	}
	//生成正常p2pkh地址
	serializedPubKey := btcwif.SerializePubKey()
	//从pubkey哈希生成正常的p2wkh地址
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, k.Opt.Params)
	//创建一个向后节点兼容的地址
	serializedScript, err := txscript.PayToAddrScript(addressWitnessPubKeyHash)
	if err != nil {
		return "", err
	}
	addressScriptHash, err := btcutil.NewAddressScriptHash(serializedScript, k.Opt.Params)
	if err != nil {
		return "", err
	}
	NativeSegWitAddress := addressScriptHash.EncodeAddress()
	return NativeSegWitAddress, nil
}

// TapRootAddr tapRoot地址 //todo 待实现
func (k *Key) TapRootAddr() (string, error) {
	km, _ := NewKeyManager(k.Opt.Mnemonic)
	key, _ := km.GetKey(config.PurposeBIP86, config.BTC, uint32(k.Opt.Account), 0, uint32(k.Opt.AddressIndex))
	prvKey, _ := btcec.PrivKeyFromBytes(key.bip32Key.Key)
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return "", err
	}
	//生成正常p2pkh地址
	serializedPubKey := btcwif.SerializePubKey()
	//从pubkey哈希生成正常的p2wkh地址
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, k.Opt.Params)
	if err != nil {
		return "", err
	}
	nestedSegWitAddr := addressWitnessPubKeyHash.EncodeAddress()
	return nestedSegWitAddr, nil
}

func (k *Key) GetPriKey(types int) (string, error) {
	Purpose := config.PurposeBIP44
	switch types {
	case 1:
		Purpose = config.PurposeBIP44
	case 2:
		Purpose = config.PurposeBIP49
	case 3:
		Purpose = config.PurposeBIP84
	default:
		Purpose = config.PurposeBIP44
	}
	km, _ := NewKeyManager(k.Opt.Mnemonic)
	key, _ := km.GetKey(Purpose, config.BTC, uint32(k.Opt.Account), 0, uint32(k.Opt.AddressIndex))
	prvKey, _ := btcec.PrivKeyFromBytes(key.bip32Key.Key)
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, true)
	return btcwif.String(), err
}

// GetKey  KeyManager 兼容一下批量创建的时候
func (km *KeyManager) GetKey(purpose, coinType, account, change, index uint32) (*BipKey, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d/%d`, config.Apostrophe, config.Apostrophe, account, change, index)

	key, ok := km.getKey(path)
	if ok {
		return &BipKey{path: path, bip32Key: key}, nil
	}

	parent, err := km.GetChangeKey(purpose, coinType, account, change)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(index)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return &BipKey{path: path, bip32Key: key}, nil
}
func (km *KeyManager) GetChangeKey(purpose, coinType, account, change uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d`, config.Apostrophe, config.Apostrophe, account, change)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetAccountKey(purpose, coinType, account)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(change)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}
func (km *KeyManager) GetAccountKey(purpose, coinType, account uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'`, config.Apostrophe, config.Apostrophe, account)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetCoinTypeKey(purpose, coinType)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(account + config.Apostrophe)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}
func (km *KeyManager) GetCoinTypeKey(purpose, coinType uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'`, config.Apostrophe, config.Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetPurposeKey(purpose)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(coinType)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}
func (km *KeyManager) GetPurposeKey(purpose uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'`, config.Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetMasterKey()
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(purpose)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}
func (km *KeyManager) GetMasterKey() (*bip32.Key, error) {
	path := "m"
	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}
	key, err := bip32.NewMasterKey(km.GetSeed())
	if err != nil {
		return nil, err
	}
	km.setKey(path, key)
	return key, nil
}
func (km *KeyManager) GetSeed() []byte {
	return bip39.NewSeed(km.GetMnemonic(), km.GetPassphrase())
}
func (km *KeyManager) getKey(path string) (*bip32.Key, bool) {
	km.mux.Lock()
	defer km.mux.Unlock()
	key, ok := km.keys[path]
	return key, ok
}
func (km *KeyManager) GetMnemonic() string {
	return km.mnemonic
}
func (km *KeyManager) GetPassphrase() string {
	return km.passphrase
}
func (km *KeyManager) setKey(path string, key *bip32.Key) {
	km.mux.Lock()
	defer km.mux.Unlock()
	km.keys[path] = key
}
