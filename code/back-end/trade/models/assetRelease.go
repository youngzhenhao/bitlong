package models

import "gorm.io/gorm"

type AssetRelease struct {
	gorm.Model
	AssetName     string `json:"asset_name" gorm:"type:varchar(255)"`
	AssetId       string `json:"asset_id" gorm:"type:varchar(255)"`
	AssetType     int    `json:"asset_type"`
	ReleaseUserId int    `json:"release_user_id"`
	ReleaseTime   int    `json:"release_time"`
	Status        int    `json:"status" gorm:"default:1"`
}
