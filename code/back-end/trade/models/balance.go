package models

import (
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	AccountId   uint    `gorm:"unique;column:account_id" json:"accountId"` // 正确地将unique和column选项放在同一个gorm标签内
	BillType    int16   `gorm:"column:bill_type;type:smallint" json:"billType"`
	Away        uint    `gorm:"column:away" json:"away"`
	Amount      float64 `gorm:"column:amount;type:decimal(10,2)" json:"amount"`
	Unit        int     `gorm:"column:Unit" json:"unit"`
	Invoice     *string `gorm:"column:invoice" json:"invoice"`
	PaymentHash *string `gorm:"column:payment_hash" json:"paymentHash"`
	Status      int16   `gorm:"column:status;type:smallint" json:"status"`
}

func (Balance) TableName() string {
	return "bill_balance"
}
