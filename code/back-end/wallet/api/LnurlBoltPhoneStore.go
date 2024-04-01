package api

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

type Invoice struct {
	ID         string `json:"id"`
	Amount     int    `json:"amount"`
	InvoiceStr string `json:"invoiceStr"`
}

type PhoneStore struct {
	DB *bolt.DB
}

func InitPhoneDB() error {
	_, err := createBucketInPhoneDB("./phone.db", "invoices")
	return err
}

func createBucketInPhoneDB(DBName, bucket string) (*bolt.Bucket, error) {
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

func (s *PhoneStore) AllInvoices(bucket string) ([]Invoice, error) {
	var invoices []Invoice
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.ForEach(func(k, v []byte) error {
			var u Invoice
			err := json.Unmarshal(v, &u)
			if err != nil {
				return err
			}
			invoices = append(invoices, u)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

// CURD

func (s *PhoneStore) CreateOrUpdateInvoice(bucket string, i *Invoice) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		buf, err := json.Marshal(i)
		if err != nil {
			return err
		}
		return b.Put([]byte(i.ID), buf)
	})
}

func (s *PhoneStore) ReadInvoice(bucket string, ID string) (*Invoice, error) {
	var i Invoice
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		userData := b.Get([]byte(ID))
		if userData == nil {
			return fmt.Errorf("no user found with ID: %s", ID)
		}
		return json.Unmarshal(userData, &i)
	})
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (s *PhoneStore) DeleteInvoice(bucket string, ID string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete([]byte(ID))
	})
}
