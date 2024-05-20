package services

import (
	"trade/dao"
	"trade/models"
)

// CreateAccount creates a new account
func CreateAccount(account *models.Account) error {
	return dao.DB.Create(account).Error
}

// GetAccount retrieves an account by ID
func GetAccount(id uint) (*models.Account, error) {
	var account models.Account
	err := dao.DB.First(&account, id).Error
	return &account, err
}

// UpdateAccount updates an existing account
func UpdateAccount(account *models.Account) error {
	return dao.DB.Save(account).Error
}

func DeleteAccount(id uint) error {
	var account models.Account
	return dao.DB.Delete(&account, id).Error
}
