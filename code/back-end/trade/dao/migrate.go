package dao

import "trade/models"

func Migrate() error {
	var err error
	if err = DB.AutoMigrate(&models.Account{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.Balance{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.BalanceExt{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.Invoice{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.FairLaunchInfo{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.FairLaunchMintedInfo{}); err != nil {
		return err
	}
	return err
}
