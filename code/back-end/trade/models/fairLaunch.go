package models

import "gorm.io/gorm"

type FairLaunchInfo struct {
	gorm.Model
	AssetID      string `json:"asset_id" gorm:"type:varchar(255);not null"`
	Name         string `json:"name" gorm:"type:varchar(255);not null"`
	Amount       int    `json:"amount"`
	Reserved     int    `json:"reserved"`
	MintQuantity int    `json:"mint_quantity"`
	StartTime    int    `json:"start_time"`
	EndTime      int    `json:"end_time"`
	Status       int    `json:"status" gorm:"default:1"`
}

type FairLaunchMintedInfo struct {
	gorm.Model
	FairLaunchInfoID int    `json:"fair_launch_info_id" gorm:"not null"`
	EncodedAddr      string `json:"encoded_addr" gorm:"type:varchar(255)"`
	AssetID          string `json:"asset_id" gorm:"type:varchar(255);not null"`
	AssetType        string `json:"asset_type" gorm:"type:varchar(255)"`
	Amount           int    `json:"amount"`
	ScriptKey        string `json:"script_key" gorm:"type:varchar(255)"`
	InternalKey      string `json:"internal_key" gorm:"type:varchar(255)"`
	TaprootOutputKey string `json:"taproot_output_key" gorm:"type:varchar(255)"`
	ProofCourierAddr string `json:"proof_courier_addr" gorm:"type:varchar(255)"`
	AssetVersion     string `json:"asset_version" gorm:"type:varchar(255)"`
	MintTime         int    `json:"mint_time"`
	Outpoint         string `json:"outpoint" gorm:"type:varchar(255)"`
	Address          string `json:"address" gorm:"type:varchar(255)"`
}
