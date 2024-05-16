package models

import "gorm.io/gorm"

type CustodyAccount struct {
	gorm.Model
	UserID          uint   `json:"user_id"`
	UserName        string `json:"user_name"`
	UserAccountCode int64  `json:"user_account_code"`
	AccountID       string `gorm:"unique" json:"account_id"`
	Label           string `json:"account_label"`
	State           uint   `json:"state"`
}

type Invoices struct {
	gorm.Model
	Invoice      string `gorm:"unique" json:"invoice"`
	AccountID    string `json:"account_id"`
	Amount       uint   `json:"amount"`
	CreationDate int64  `json:"creation_date"`
	Expiry       int64  `json:"expiry"`
	Status       string `json:"status"`
	Settled      bool   `json:"settled"`
}
