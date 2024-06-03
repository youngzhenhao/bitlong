package api

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

type ServerStore struct {
	DB *bolt.DB
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Socket     string `json:"socket"`
	RemotePort string `json:"remote_port"`
}

func InitServerDB() error {
	_, err := createBucketInServerDB("./server.db", "users")
	return err
}

func createBucketInServerDB(DBName, bucket string) (*bolt.Bucket, error) {
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

func (s *ServerStore) AllUsers(bucket string) ([]User, error) {
	var users []User
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.ForEach(func(k, v []byte) error {
			var u User
			err := json.Unmarshal(v, &u)
			if err != nil {
				return err
			}
			users = append(users, u)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CURD

func (s *ServerStore) CreateOrUpdateUser(bucket string, u *User) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}
		return b.Put([]byte(u.ID), buf)
	})
}

func (s *ServerStore) ReadUser(bucket string, ID string) (*User, error) {
	var u User
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		userData := b.Get([]byte(ID))
		if userData == nil {
			return fmt.Errorf("no user found with ID: %s", ID)
		}
		return json.Unmarshal(userData, &u)
	})
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *ServerStore) DeleteUser(bucket string, ID string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete([]byte(ID))
	})
}
