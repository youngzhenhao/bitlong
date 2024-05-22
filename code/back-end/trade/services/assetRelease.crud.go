package services

import (
	"gorm.io/gorm"
	"trade/models"
)

type AssetReleaseStore struct {
	DB *gorm.DB
}

// AssetRelease

func (a *AssetReleaseStore) CreateAssetRelease(assetRelease *models.AssetRelease) error {
	return a.DB.Create(assetRelease).Error
}

func (a *AssetReleaseStore) ReadAssetRelease(id uint) (*models.AssetRelease, error) {
	var assetRelease models.AssetRelease
	err := a.DB.First(&assetRelease, id).Error
	return &assetRelease, err
}

func (a *AssetReleaseStore) UpdateAssetRelease(assetRelease *models.AssetRelease) error {
	return a.DB.Save(assetRelease).Error
}

func (a *AssetReleaseStore) DeleteAssetRelease(id uint) error {
	var assetRelease models.AssetRelease
	return a.DB.Delete(&assetRelease, id).Error
}
