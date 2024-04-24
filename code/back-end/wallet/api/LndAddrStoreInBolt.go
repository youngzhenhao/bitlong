package api

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/wallet/base"
	"path/filepath"
	"time"
)

type Addr struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
	Type    string `json:"type"`
	Path    string `json:"path"`
}

type AddrStore struct {
	DB *bolt.DB
}

func InitAddrDB() error {
	_, err := createBucketInAddrDB(filepath.Join(base.QueryConfigByKey("dirpath"), "addr.db"), "addresses")
	return err
}

func createBucketInAddrDB(DBName, bucket string) (*bolt.Bucket, error) {
	db, err := bolt.Open(DBName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
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

func (s *AddrStore) AllAddresses(bucket string) ([]Addr, error) {
	var Addrs []Addr
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.ForEach(func(k, v []byte) error {
			var u Addr
			err := json.Unmarshal(v, &u)
			if err != nil {
				return err
			}
			Addrs = append(Addrs, u)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return Addrs, nil
}

// CURD

func (s *AddrStore) CreateOrUpdateAddr(bucket string, a *Addr) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		buf, err := json.Marshal(a)
		if err != nil {
			return err
		}
		return b.Put([]byte(a.Address), buf)
	})
}

func (s *AddrStore) ReadAddr(bucket string, address string) (*Addr, error) {
	var a Addr
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		userData := b.Get([]byte(address))
		if userData == nil {
			return fmt.Errorf("no user found with address: %s", address)
		}
		return json.Unmarshal(userData, &a)
	})
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (s *AddrStore) DeleteAddr(bucket string, address string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete([]byte(address))
	})
}
