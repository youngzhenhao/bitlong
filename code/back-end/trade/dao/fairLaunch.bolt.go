package dao

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"time"
	"trade/config"
	"trade/models"
	"trade/utils"
)

var (
	serverDbPath          = config.GetLoadConfig().Bolt.DbPath
	serverDbMode          = config.GetLoadConfig().Bolt.DbMode
	serverDbTimeoutSecond = config.GetLoadConfig().Bolt.DbTimeoutSecond
	fairLaunchBucketName  = "fair_launch"
)

func GetServerDbPath() string {
	if serverDbPath == "" {
		return "/root/database/server.db"
	}
	return serverDbPath
}

func GetServerDbMode() os.FileMode {
	if serverDbMode == 0 {
		return 0600
	}
	return serverDbMode
}

func GetServerDbTimeout() time.Duration {
	return getServerDbTimeoutSecond() * time.Second
}

func getServerDbTimeoutSecond() time.Duration {
	if serverDbTimeoutSecond == 0 {
		return 1
	}
	return serverDbTimeoutSecond
}

func GetFairLaunchBucketName() string {
	return fairLaunchBucketName
}

type ServerStore struct {
	DB *bolt.DB `json:"db"`
}

func InitServerDB() error {
	_, err := createBucketInServerDB(GetServerDbPath(), GetFairLaunchBucketName())
	return err
}

func createBucketInServerDB(DBName, bucket string) (*bolt.Bucket, error) {
	db, err := bolt.Open(DBName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		utils.LogError("bolt.Open", err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			utils.LogError("db.Close", err)
		}
	}(db)
	var b *bolt.Bucket
	err = db.Update(func(tx *bolt.Tx) error {
		b, err = tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *ServerStore) AllFairLaunchInfo(bucket string) (*[]models.FairLaunchInfo, error) {
	var fairLaunchInfos []models.FairLaunchInfo
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.ForEach(func(k, v []byte) error {
			var f models.FairLaunchInfo
			err := json.Unmarshal(v, &f)
			if err != nil {
				return err
			}
			fairLaunchInfos = append(fairLaunchInfos, f)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return &fairLaunchInfos, nil
}

// CURD

func (s *ServerStore) CreateOrUpdateFairLaunchInfo(bucket string, f *models.FairLaunchInfo) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		buf, err := json.Marshal(f)
		if err != nil {
			return err
		}
		return b.Put([]byte(f.ID), buf)
	})
}

func (s *ServerStore) ReadFairLaunchInfo(bucket string, ID string) (*models.FairLaunchInfo, error) {
	var f models.FairLaunchInfo
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		fairLaunchInfo := b.Get([]byte(ID))
		if fairLaunchInfo == nil {
			return fmt.Errorf("no FairLaunchInfo found with ID: %s", ID)
		}
		return json.Unmarshal(fairLaunchInfo, &f)
	})
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (s *ServerStore) DeleteFairLaunchInfo(bucket string, ID string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete([]byte(ID))
	})
}
