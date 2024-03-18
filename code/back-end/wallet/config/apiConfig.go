package config

import (
	"github.com/btcsuite/btcd/chaincfg"
)

var (
	BTCParams = chaincfg.MainNetParams
)

// zero is deafult of uint32
const (
	PrvType          = "Hex"
	Zero      uint32 = 0
	ZeroQuote uint32 = 0x80000000
	BTCToken  uint32 = 0x10000000

	PurposeBIP44 uint32 = 0x8000002C // 44' BIP44
	PurposeBIP49 uint32 = 0x80000031 // 49' BIP49
	PurposeBIP84 uint32 = 0x80000054 // 84' BIP84
	PurposeBIP86 uint32 = 0x80000056 // 86' BIP86
	Apostrophe   uint32 = 0x80000000

	//交易状态
	TransactionStatusOk      = "success"
	TransactionStatusFaild   = "fail"
	TransactionStatusPending = "pending"
)

// bip44钱包类型  16进制
const (
	// https://github.com/satoshilabs/slips/blob/master/slip-0044.md#registered-coin-types
	BTC = ZeroQuote + 0
	// btc token
	USDT = BTCToken + 1
)

// BlockDamon
var (
	BlockDaemonApiKey       = "yK-9xGBXPIvgc0YGBqrHv6l7JR7P8M-BeGUBPA9nN6cFfZMe"
	BlockDamonApiUrl        = "https://svc.blockdaemon.com/"
	BlockDamonGetBalanceUrl = "universal/v1/%s/%s/account/%s?apiKey=" + BlockDaemonApiKey
	BlockDamonGetUTXOUrl    = "universal/v1/%s/%s/account/%s/utxo?apiKey=" + BlockDaemonApiKey + "&page_size=100"
	BlockDamonGetGasFeeUrl  = "universal/v1/%s/%s/tx/estimate_fee?apiKey=" + BlockDaemonApiKey
	BlockDamonTxSendUrl     = "universal/v1/%s/%s/tx/send?apiKey=" + BlockDaemonApiKey
	BlockDamonTxListUrl     = "universal/v1/%s/%s/account/%s/txs?apiKey=" + BlockDaemonApiKey + "&page_token=%s&page_size=%d"
)
