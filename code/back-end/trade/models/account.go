package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserId          uint    `gorm:"column:user_id" json:"userId"` // column选项放在同一个gorm标签内
	UserName        string  `gorm:"column:user_name" json:"userName"`
	UserAccountCode string  `gorm:"column:user_account_code" json:"userAccountCode"`
	Status          int16   `gorm:"column:status;type:smallint" json:"status"`
	Label           *string `gorm:"column:label" json:"label"`
}

func (Account) TableName() string {
	return "user_account"
}
