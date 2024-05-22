package models

import "gorm.io/gorm"

type AssetRelease struct {
	gorm.Model
	AssetName     string `json:"asset_name" gorm:"type:varchar(255)"`
	AssetId       string `json:"asset_id" gorm:"type:varchar(255)"`
	AssetType     string `json:"asset_type" gorm:"type:varchar(255)"`
	ReleaseUserId int    `json:"release_user_id"`
	ReleaseTime   int    `json:"release_time"`
	State         int    `json:"state" gorm:"default:0"`
}
