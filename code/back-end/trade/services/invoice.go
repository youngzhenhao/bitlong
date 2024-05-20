package services

import (
	"gorm.io/gorm"
	"trade/dao"
	"trade/models"
)

// CreateInvoice creates a new invoice record
func CreateInvoice(invoice *models.Invoice) error {
	return dao.DB.Create(invoice).Error
}

// GetInvoice retrieves an invoice by ID
func GetInvoice(id uint) (*models.Invoice, error) {
	var invoice models.Invoice
	err := dao.DB.First(&invoice, id).Error
	return &invoice, err
}

// UpdateInvoice updates an existing invoice
func UpdateInvoice(db *gorm.DB, invoice *models.Invoice) error {
	return db.Save(invoice).Error
}

// DeleteInvoice soft deletes an invoice by ID
func DeleteInvoice(id uint) error {
	var invoice models.Invoice
	return dao.DB.Delete(&invoice, id).Error
}
