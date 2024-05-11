package dao

import (
	"AssetsTrade/api"
	"AssetsTrade/models"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

type ServerStore struct {
	DB *bolt.DB
}

func InitServerDB() error {
	_, err := createBucketInServerDB("/root/database/server.db", "fair_launch")
	return err
}

func createBucketInServerDB(DBName, bucket string) (*bolt.Bucket, error) {
	db, err := bolt.Open(DBName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", api.GetTimeNow(), err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", api.GetTimeNow(), err)
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
