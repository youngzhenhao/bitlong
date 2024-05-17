package models

import (
	"gorm.io/gorm"
	"time"
)

type Invoice struct {
	gorm.Model
	UserID     uint       `gorm:"not null;column:user_id" json:"userId"`
	AccountID  *uint      `gorm:"column:account_id" json:"accountId"`
	Amount     float64    `gorm:"type:decimal(10,2);column:amount" json:"amount"`
	CreateDate *time.Time `gorm:"column:create_date" json:"createDate"`
	Expiry     *int       `gorm:"column:expiry" json:"expiry"`
	Status     int16      `gorm:"column:status;type:smallint" json:"status"`
}

func (Invoice) TableName() string {
	return "user_invoice"
}
