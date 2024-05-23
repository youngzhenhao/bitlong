package models

import (
	"gorm.io/gorm"
)

var (
	StatusNormal     = 1
	StatusPending    = 2
	StatusDeprecated = 0
)

type FairLaunchInfo struct {
	gorm.Model
	ImageData              string  `json:"image_data"`
	Name                   string  `json:"name" gorm:"type:varchar(255);not null"`
	AssetType              int     `json:"asset_type"`
	Amount                 int     `json:"amount"`
	Reserved               int     `json:"reserved"`
	MintQuantity           int     `json:"mint_quantity"`
	StartTime              int     `json:"start_time"`
	EndTime                int     `json:"end_time"`
	Description            string  `json:"description"`
	ActualReserved         float64 `json:"actual_reserved"`
	ReserveTotal           int     `json:"reserve_total"`
	MintNumber             int     `json:"mint_number"`
	IsFinalEnough          bool    `json:"is_final_enough"`
	FinalQuantity          int     `json:"final_quantity"`
	MintTotal              int     `json:"mint_total"`
	ActualMintTotalPercent float64 `json:"actual_mint_total_percent"`
	CalculationExpression  string  `json:"calculation_expression" gorm:"type:varchar(255)"`
	BatchKey               string  `json:"batch_key" gorm:"type:varchar(255)"`
	BatchState             string  `json:"batch_state" gorm:"type:varchar(255)"`
	BatchTxidAnchor        string  `json:"batch_txid_anchor" gorm:"type:varchar(255)"`
	AssetID                string  `json:"asset_id" gorm:"type:varchar(255)"`
	UserID                 int     `json:"user_id"`
	IsReservedSent         bool    `json:"is_reserved_sent"`
	MintedNumber           int     `json:"minted_number"`
	IsMintAll              bool    `json:"is_mint_all"`
	Status                 int     `json:"status" default:"1" gorm:"default:1"`
}

type FairLaunchMintedInfo struct {
	gorm.Model
	FairLaunchInfoID int    `json:"fair_launch_info_id" gorm:"not null"`
	EncodedAddr      string `json:"encoded_addr" gorm:"type:varchar(255)"`
	AssetID          string `json:"asset_id" gorm:"type:varchar(255)"`
	AssetType        string `json:"asset_type" gorm:"type:varchar(255)"`
	AddrAmount       int    `json:"amount_addr"`
	ScriptKey        string `json:"script_key" gorm:"type:varchar(255)"`
	InternalKey      string `json:"internal_key" gorm:"type:varchar(255)"`
	TaprootOutputKey string `json:"taproot_output_key" gorm:"type:varchar(255)"`
	ProofCourierAddr string `json:"proof_courier_addr" gorm:"type:varchar(255)"`
	MintTime         int    `json:"mint_time"`
	Outpoint         string `json:"outpoint" gorm:"type:varchar(255)"`
	Address          string `json:"address" gorm:"type:varchar(255)"`
	Status           int    `json:"status" gorm:"default:1"`
}

type FairLaunchInventoryInfo struct {
	gorm.Model
	FairLaunchInfoID int  `json:"fair_launch_info_id" gorm:"not null"`
	Quantity         int  `json:"quantity"`
	IsMinted         bool `json:"is_minted"`
	MintedID         int  `json:"minted_id"`
	Status           int  `json:"status" gorm:"default:1"`
}
