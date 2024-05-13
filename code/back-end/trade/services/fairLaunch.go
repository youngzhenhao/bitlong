package services

import (
	"github.com/boltdb/bolt"
	"trade/dao"
	"trade/models"
	"trade/utils"
)

func GetFairLaunch(id string) *models.FairLaunchInfo {
	_ = dao.InitServerDB()
	db, err := bolt.Open(dao.GetServerDbPath(), dao.GetServerDbMode(), &bolt.Options{Timeout: dao.GetServerDbTimeout()})
	if err != nil {
		utils.LogError("bolt.Open", err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			utils.LogError("db.Close", err)
		}
	}(db)
	s := &dao.ServerStore{DB: db}
	fairLaunchInfo, err := s.ReadFairLaunchInfo(dao.GetFairLaunchBucketName(), id)
	if err != nil {
		return nil
	}
	return fairLaunchInfo
}

func GetMinted(id string) *[]models.MintedInfo {
	fairLaunchInfo := GetFairLaunch(id)
	if fairLaunchInfo == nil || fairLaunchInfo.Minted == nil {
		return nil
	}
	return fairLaunchInfo.Minted
}

func SetFairLaunch() {
	_ = dao.InitServerDB()
	db, err := bolt.Open(dao.GetServerDbPath(), dao.GetServerDbMode(), &bolt.Options{Timeout: dao.GetServerDbTimeout()})
	if err != nil {
		utils.LogError("bolt.Open", err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			utils.LogError("db.Close", err)
		}
	}(db)
	s := &dao.ServerStore{DB: db}
	var mintedInfo []models.MintedInfo
	mintedInfo = append(mintedInfo, models.MintedInfo{})
	err = s.CreateOrUpdateFairLaunchInfo("fair_launch", &models.FairLaunchInfo{})
	if err != nil {
		return
	}
	// TODO: need to complete
}

func FairLaunchMint() {
	// TODO: need to complete
	// @dev: Call dao's api

}

func GetAllFairLaunch() *[]models.FairLaunchInfo {
	_ = dao.InitServerDB()
	db, err := bolt.Open(dao.GetServerDbPath(), dao.GetServerDbMode(), &bolt.Options{Timeout: dao.GetServerDbTimeout()})
	if err != nil {
		utils.LogError("bolt.Open", err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			utils.LogError("db.Close", err)
		}
	}(db)
	s := &dao.ServerStore{DB: db}
	fairLaunchInfos, err := s.AllFairLaunchInfo(dao.GetFairLaunchBucketName())
	if err != nil {
		return nil
	}
	return fairLaunchInfos
}
