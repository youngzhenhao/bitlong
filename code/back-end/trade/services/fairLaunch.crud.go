package services

import (
	"gorm.io/gorm"
	"trade/models"
)

type FairLaunchStore struct {
	DB *gorm.DB
}

// FairLaunchInfo

func (f *FairLaunchStore) CreateFairLaunchInfo(fairLaunchInfo *models.FairLaunchInfo) error {
	return f.DB.Create(fairLaunchInfo).Error
}

func (f *FairLaunchStore) ReadFairLaunchInfo(id uint) (*models.FairLaunchInfo, error) {
	var fairLaunchInfo models.FairLaunchInfo
	err := f.DB.First(&fairLaunchInfo, id).Error
	return &fairLaunchInfo, err
}

func (f *FairLaunchStore) UpdateFairLaunchInfo(fairLaunchInfo *models.FairLaunchInfo) error {
	return f.DB.Save(fairLaunchInfo).Error
}

func (f *FairLaunchStore) DeleteFairLaunchInfo(id uint) error {
	var fairLaunchInfo models.FairLaunchInfo
	return f.DB.Delete(&fairLaunchInfo, id).Error
}

// FairLaunchMintedInfo

func (f *FairLaunchStore) CreateFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) error {
	return f.DB.Create(fairLaunchMintedInfo).Error
}

func (f *FairLaunchStore) ReadFairLaunchMintedInfo(id uint) (*models.FairLaunchMintedInfo, error) {
	var fairLaunchMintedInfo models.FairLaunchMintedInfo
	err := f.DB.First(&fairLaunchMintedInfo, id).Error
	return &fairLaunchMintedInfo, err
}

func (f *FairLaunchStore) UpdateFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) error {
	return f.DB.Save(fairLaunchMintedInfo).Error
}

func (f *FairLaunchStore) DeleteFairLaunchMintedInfo(id uint) error {
	var fairLaunchMintedInfo models.FairLaunchMintedInfo
	return f.DB.Delete(&fairLaunchMintedInfo, id).Error
}
