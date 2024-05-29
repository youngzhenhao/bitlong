package models

import "gorm.io/gorm"

type (
	AssetIssuanceState int
)

var (
	AssetIssuanceStatePending AssetIssuanceState = 0
	AssetIssuanceStateIssued  AssetIssuanceState = 1
)

type AssetIssuance struct {
	gorm.Model
	AssetName      string             `json:"asset_name" gorm:"type:varchar(255)"`
	AssetId        string             `json:"asset_id" gorm:"type:varchar(255)"`
	AssetType      int                `json:"asset_type"`
	IssuanceUserId int                `json:"issuance_user_id"`
	IssuanceTime   int                `json:"issuance_time"`
	IsFairLaunch   bool               `json:"is_fair_launch"`
	FairLaunchID   int                `json:"fair_launch_id"`
	Status         int                `json:"status" gorm:"default:1"`
	State          AssetIssuanceState `json:"state"`
}
