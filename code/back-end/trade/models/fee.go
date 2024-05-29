package models

import "gorm.io/gorm"

type FeeRateInfo struct {
	gorm.Model
	Name                 string  `json:"name" gorm:"type:varchar(255);not null"`
	EstimateSmartFeeRate float64 `json:"estimate_smart_fee_rate"`
}
