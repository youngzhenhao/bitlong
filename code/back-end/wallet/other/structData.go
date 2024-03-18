package other

import "encoding/json"

type TransactionBase struct {
	TxHash    string `json:"txHash"`              // 事物hash
	Chain     string `json:"chain"`               // 链名称
	Height    string `json:"height"`              // 区块高度
	Time      string `json:"time,omitempty"`      // 交易时间
	Status    string `json:"status,omitempty"`    // 状态
	From      string `json:"from,omitempty"`      // 转账地址
	To        string `json:"to,omitempty"`        // 目标地址
	Fee       string `json:"fee,omitempty"`       // 交易手续费
	FeeCny    string `json:"feeCny,omitempty"`    // 交易手续费
	FeeUsd    string `json:"feeUsd,omitempty"`    // 交易手续费
	Confirm   string `json:"confirm,omitempty"`   // 交易确认数
	Amount    string `json:"amount,omitempty"`    // 交易数量
	TokenName string `json:"tokenName,omitempty"` // 代币名称
	Send      bool   `json:"send"`                // 是否是发送
	//PageToken  string        `json:"-"`
	UtxoInput  []*UtxoAmount `json:"utxoInput,omitempty"`  // btc链存在多个输入和多个输出，btc等utxo模型链使用
	UtxoOutput []*UtxoAmount `json:"utxoOutput,omitempty"` // btc链存在多个输入和多个输出，btc等utxo模型链使用
}

type TransactionList struct {
	List      []*TransactionBase `json:"list"`
	PageToken string             `json:"next_page_token"`
	TotalPage string             `json:"total_page"`
}

type TransactionDetail struct {
	TransactionBase
}

type UtxoAmount struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type TxListData struct {
	Data []struct {
		Id            string      `json:"id"`
		Date          int         `json:"date"`
		Status        string      `json:"status"`
		Confirmations int         `json:"confirmations"`
		BlockNumber   json.Number `json:"block_number"`
		Events        []EventData `json:"events"`
	} `json:"data"`
	Meta struct {
		Paging struct {
			NextPageToken string `json:"next_page_token"`
		} `json:"paging"`
	} `json:"meta"`
}

type TxListData4json struct {
	Data []struct {
		Id            string           `json:"id"`
		Date          int              `json:"date"`
		Status        string           `json:"status"`
		Confirmations int              `json:"confirmations"`
		BlockNumber   json.Number      `json:"block_number"`
		Events        []EventData4Json `json:"events"`
	} `json:"data"`
	Meta struct {
		Paging struct {
			NextPageToken string `json:"next_page_token"`
		} `json:"paging"`
	} `json:"meta"`
}

// BalanceRegData 余额数据集
type BalanceRegData struct {
	Currency struct {
		AssetPath string `json:"asset_path"`
		Symbol    string `json:"symbol"`
		Name      string `json:"name"`
		Decimals  int    `json:"decimals"`
		Type      string `json:"type"`
	} `json:"currency"`
	ConfirmedBalance string `json:"confirmed_balance"`
	ConfirmedBlock   int    `json:"confirmed_block"`
}
type EventData struct {
	Type         string      `json:"type"`
	Denomination string      `json:"denomination"`
	Decimals     json.Number `json:"decimals"`
	From         string      `json:"source"`
	Amount       float64     `json:"amount"`
	To           string      `json:"destination,omitempty"`
}

type EventData4Json struct {
	Type         string      `json:"type"`
	Denomination string      `json:"denomination"`
	Decimals     json.Number `json:"decimals"`
	From         string      `json:"source"`
	Amount       float64     `json:"amount"`
	To           string      `json:"destination,omitempty"`
}

// GasFeeData 三种手续费 快-推荐-慢
type GasFeeData struct {
	MostRecentBlock int `json:"most_recent_block"`
	EstimatedFees   struct {
		Fast   int `json:"fast"`
		Medium int `json:"medium"`
		Slow   int `json:"slow"`
	} `json:"estimated_fees"`
}

type SendTxRegData struct {
	TxId string `json:"id"`
}

// UTXORegData utxo数据集
type UTXORegData struct {
	Total int        `json:"total"`
	Data  []UTXOdata `json:"data"`
	Meta  struct {
		Paging struct {
			NextPageToken string `json:"next_page_token"`
		} `json:"paging"`
	} `json:"meta"`
}

type UTXOdata struct {
	Status  string `json:"status"`
	IsSpent bool   `json:"is_spent"`
	Value   int64  `json:"value"`
	Mined   struct {
		Index         int    `json:"index"`
		TxId          string `json:"tx_id"`
		Date          int    `json:"date"`
		BlockId       string `json:"block_id"`
		BlockNumber   int    `json:"block_number"`
		Confirmations int    `json:"confirmations"`
		Meta          struct {
			Index      int      `json:"index"`
			Script     string   `json:"script"`
			Addresses  []string `json:"addresses"`
			ScriptType string   `json:"script_type"`
		} `json:"meta"`
	} `json:"mined"`
	Spent struct {
		TxId          string `json:"tx_id"`
		Date          int    `json:"date"`
		BlockId       string `json:"block_id"`
		BlockNumber   int    `json:"block_number"`
		Confirmations int    `json:"confirmations"`
	} `json:"spent,omitempty"`
}
