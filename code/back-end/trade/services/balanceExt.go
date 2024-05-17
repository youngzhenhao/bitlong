package services

import (
	"trade/middleware"
	"trade/models"
)

// CreateBalanceExt creates a new balance extension record
func CreateBalanceExt(balanceExt *models.BalanceExt) error {
	return middleware.DB.Create(balanceExt).Error
}

// GetBalanceExt retrieves a balance extension by ID
func GetBalanceExt(id uint) (*models.BalanceExt, error) {
	var balanceExt models.BalanceExt
	err := middleware.DB.First(&balanceExt, id).Error
	return &balanceExt, err
}

// UpdateBalanceExt updates an existing balance extension
func UpdateBalanceExt(balanceExt *models.BalanceExt) error {
	return middleware.DB.Save(balanceExt).Error
}

// DeleteBalanceExt soft deletes a balance extension by ID
func DeleteBalanceExt(id uint) error {
	var balanceExt models.BalanceExt
	return middleware.DB.Delete(&balanceExt, id).Error
}
