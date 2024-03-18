package address

import (
	"crypto/ecdsa"
	"crypto/sha512"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/vedhavyas/go-subkey"
	"golang.org/x/crypto/pbkdf2"
)

type GetKeys struct {
	Mnemonic     string
	CoinType     int64
	Account      int64
	Purpose      int64
	AddressIndex int64
}
type Key struct {
	Opt      *Options
	Extended *hdkeychain.ExtendedKey

	// for btc
	Private *btcec.PrivateKey
	Public  *btcec.PublicKey

	// for eth
	PrivateECDSA *ecdsa.PrivateKey
	PublicECDSA  *ecdsa.PublicKey

	// for dot
	KeyPair subkey.KeyPair
}

func (k *Key) init() error {
	var err error

	k.Private, err = k.Extended.ECPrivKey()
	if err != nil {
		return err
	}

	k.Public, err = k.Extended.ECPubKey()
	if err != nil {
		return err
	}

	k.PrivateECDSA = k.Private.ToECDSA()
	k.PublicECDSA = &k.PrivateECDSA.PublicKey

	return nil
}
func (k *Key) GetChildKey(opts ...Option) (*Key, error) {
	var (
		err error
		o   = NewOpt(opts...)
		no  = o
	)
	typ, ok := coinTypes[o.CoinType]

	if ok {
		no = NewOpt(append(opts, CoinType(typ))...)
	}
	extended := k.Extended
	fmt.Println("1", extended.String())
	fmt.Println("path", no.GetPath())
	for _, i := range no.GetPath() {
		extended, err = extended.Derive(uint32(i))
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("2", extended.String())
	key := &Key{
		Opt:      o,
		Extended: extended,
	}
	err = key.init()
	if err != nil {
		fmt.Println("GetChildKey error", err)
		return nil, err
	}
	return key, nil
}

// NewOpt 初始化opts
func NewOpt(opts ...Option) *Options {
	opt := &Options{
		Params:       DefaultParams,
		Password:     DefaultPassword,
		Language:     DefaultLanguage,
		Purpose:      DefaultPurpose,
		CoinType:     DefaultCoinType,
		Change:       DefaultChange,
		Account:      DefaultAccount,
		AddressIndex: DefaultAddressIndex,
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// HdNewKey 子地址信息
func HdNewKey(getK *GetKeys) (*Key, error) {
	var (
		err error
		o   = NewOpt(
			Mnemonic(getK.Mnemonic),         //注入助记词
			CoinType(getK.CoinType),         //注入币种
			Account(getK.Account),           //注入账户
			Purpose(getK.Purpose),           //注入账户
			AddressIndex(getK.AddressIndex), //注入地址索引
		)
	)

	if len(o.Seed) <= 0 {
		o.Seed = NewSeed(o.Mnemonic, o.Password)
	}

	if err != nil {
		return nil, err
	}

	extended, err := hdkeychain.NewMaster(o.Seed, o.Params)
	if err != nil {
		return nil, err
	}
	key := &Key{
		Opt:      o,
		Extended: extended,
	}
	err = key.init()
	if err != nil {

		fmt.Println("HdNewKey error", err)
		return nil, err
	}
	return key, nil
}

// 还原seed
func NewSeed(mnemonic string, password string) []byte {
	return pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+password), 2048, 64, sha512.New)
}
