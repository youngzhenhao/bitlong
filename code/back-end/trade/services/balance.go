package services

import (
	"trade/dao"
	"trade/models"
)

// CreateBalance creates a new balance record
func CreateBalance(balance *models.Balance) error {
	return dao.DB.Create(balance).Error
}

// GetBalance retrieves a balance by ID
func GetBalance(id uint) (*models.Balance, error) {
	var balance models.Balance
	err := dao.DB.First(&balance, id).Error
	return &balance, err
}

// UpdateBalance updates an existing balance
func UpdateBalance(balance *models.Balance) error {
	return dao.DB.Save(balance).Error
}

// DeleteBalance soft deletes a balance by ID
func DeleteBalance(id uint) error {
	var balance models.Balance
	return dao.DB.Delete(&balance, id).Error
}
