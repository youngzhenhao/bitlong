package address

import (
	"github.com/wallet/config"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
)

var coinTypes = map[int64]int64{
	int64(config.USDT): int64(config.BTC),
}

// default options
var (
	DefaultParams = &config.BTCParams

	// master key options
	DefaultPassword = ""
	DefaultLanguage = ""

	// child key options
	DefaultPurpose      = int64(config.ZeroQuote) + 44
	DefaultCoinType     = int64(config.BTC)
	DefaultAccount      = int64(config.ZeroQuote)
	DefaultChange       = int64(config.Zero)
	DefaultAddressIndex = int64(config.Zero)
)

// Option of key
type Option func(*Options)

// Options of key
type Options struct {
	Params *chaincfg.Params

	// master key options
	Mnemonic string
	Password string
	Language string
	Seed     []byte

	// child key options
	Purpose      int64
	CoinType     int64
	Account      int64
	Change       int64
	AddressIndex int64
}

// GetPath return path in bip44 style
func (o *Options) GetPath() []int64 {
	return []int64{
		o.Purpose,
		o.CoinType,
		o.Account,
		o.Change,
		o.AddressIndex,
	}
}

// Params set to options
func Params(p *chaincfg.Params) Option {
	return func(o *Options) {
		o.Params = p
	}
}

// Mnemonic set to options
func Mnemonic(m string) Option {
	return func(o *Options) {
		o.Mnemonic = m
	}
}

// Password set to options
func Password(p string) Option {
	return func(o *Options) {
		o.Password = p
	}
}

// Language set to options
func Language(l string) Option {
	return func(o *Options) {
		o.Language = l
	}
}

// Seed set to options
func Seed(s []byte) Option {
	return func(o *Options) {
		o.Seed = s
	}
}

// Purpose set to options
func Purpose(p int64) Option {
	return func(o *Options) {
		o.Purpose = p
	}
}

// CoinType set to options
func CoinType(c int64) Option {
	return func(o *Options) {
		o.CoinType = c
	}
}

// Account set to options
func Account(a int64) Option {
	return func(o *Options) {
		o.Account = a
	}
}

// Change set to options
func Change(c int64) Option {
	return func(o *Options) {
		o.Change = c
	}
}

// AddressIndex set to options
func AddressIndex(a int64) Option {
	return func(o *Options) {
		o.AddressIndex = a
	}
}

// Path set to options
// example: m/44'/0'/0'/0/0
// example: m/Purpose'/CoinType'/Account'/Change/AddressIndex
func Path(path string) Option {
	return func(o *Options) {
		path = strings.TrimPrefix(path, "m/")
		paths := strings.Split(path, "/")
		if len(paths) != 5 {
			return
		}
		o.Purpose = PathNumber(paths[0])
		o.CoinType = PathNumber(paths[1])
		o.Account = PathNumber(paths[2])
		o.Change = PathNumber(paths[3])
		o.AddressIndex = PathNumber(paths[4])
	}
}

// PathNumber 44' => 0x80000000 + 44
func PathNumber(str string) int64 {
	num64, _ := strconv.ParseInt(strings.TrimSuffix(str, "'"), 10, 64)
	num := uint32(num64)
	if strings.HasSuffix(str, "'") {
		num += config.ZeroQuote
	}
	return int64(num)
}
