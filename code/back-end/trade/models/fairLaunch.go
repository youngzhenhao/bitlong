package models

import "gorm.io/gorm"

type FairLaunchInfo struct {
	gorm.Model
	AssetID      string `json:"asset_id" gorm:"not null"`
	Name         string `json:"name"`
	Amount       int    `json:"amount"`
	Reserved     int    `json:"reserved"`
	MintQuantity int    `json:"mint_quantity"`
	StartTime    int    `json:"start_time"`
	EndTime      int    `json:"end_time"`
	Status       int    `json:"status"`
}

type FairLaunchMintedInfo struct {
	gorm.Model
	FairLaunchInfoID string `json:"fair_launch_info_id" gorm:"not null"`
	EncodedAddr      string `json:"encoded_addr"`
	AssetID          string `json:"asset_id" gorm:"not null"`
	AssetType        string `json:"asset_type"`
	Amount           int    `json:"amount"`
	ScriptKey        string `json:"script_key"`
	InternalKey      string `json:"internal_key"`
	TaprootOutputKey string `json:"taproot_output_key"`
	ProofCourierAddr string `json:"proof_courier_addr"`
	AssetVersion     string `json:"asset_version"`
	MintTime         int    `json:"mint_time"`
	Outpoint         string `json:"outpoint"`
	Address          string `json:"address"`
}
