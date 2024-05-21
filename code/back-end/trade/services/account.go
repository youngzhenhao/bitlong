package services

import (
	"trade/middleware"
	"trade/models"
)

// CreateAccount creates a new account
func CreateAccount(account *models.Account) error {
	return middleware.DB.Create(account).Error
}

// GetAccount retrieves an account by ID
func GetAccount(id uint) (*models.Account, error) {
	var account models.Account
	err := middleware.DB.First(&account, id).Error
	return &account, err
}

// ReadAccountByUserId retrieves an account by user ID
func ReadAccountByUserId(userId uint) (*models.Account, error) {
	var account models.Account
	err := middleware.DB.Where("user_id =?", userId).Find(&account).Error
	return &account, err
}

// ReadAccountByUserIds retrieves accounts by user IDs
func ReadAccountByUserIds(userId uint) ([]models.Account, error) {
	var accounts []models.Account
	err := middleware.DB.Where("user_id =?", userId).Find(&accounts).Error
	return accounts, err
}

// UpdateAccount updates an existing account
func UpdateAccount(account *models.Account) error {
	return middleware.DB.Save(account).Error
}

func DeleteAccount(id uint) error {
	var account models.Account
	return middleware.DB.Delete(&account, id).Error
}
