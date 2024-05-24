package models

import "gorm.io/gorm"

type AssetIssuance struct {
	gorm.Model
	AssetName      string `json:"asset_name" gorm:"type:varchar(255)"`
	AssetId        string `json:"asset_id" gorm:"type:varchar(255)"`
	AssetType      int    `json:"asset_type"`
	IssuanceUserId int    `json:"issuance_user_id"`
	IssuanceTime   int    `json:"issuance_time"`
	Status         int    `json:"status" gorm:"default:1"`
}
