package clientdb

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.etcd.io/bbolt"
)

const (
	// DBFilename is the default filename of the client database.
	DBFilename = "pool.db"

	// dbFilePermission is the default permission the client database file
	// is created with.
	dbFilePermission = 0600

	// DefaultPoolDBTimeout is the default maximum time we wait for the
	// Pool bbolt database to be opened. If the database is already opened
	// by another process, the unique lock cannot be obtained. With the
	// timeout we error out after the given time instead of just blocking
	// for forever.
	DefaultPoolDBTimeout = 5 * time.Second
)

var (
	// byteOrder is the default byte order we'll use for serialization
	// within the database.
	byteOrder = binary.BigEndian
)

// DB is a bolt-backed persistent store.
type DB struct {
	*bbolt.DB
}

// New creates a new bolt database that can be found at the given directory.
func New(dir, fileName string) (*DB, error) {
	firstInit := false
	path := filepath.Join(dir, fileName)

	// If the database file does not exist yet, create its directory.
	if !fileExists(path) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return nil, err
		}
		firstInit = true
	}

	db, err := initDB(path, firstInit)
	if err != nil {
		return nil, err
	}

	// Attempt to sync the database's current version with the latest known
	// version available.
	if err := syncVersions(db); err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}

// fileExists reports whether the named file or directory exists.
func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// initDB initializes all of the required top-level buckets for the database.
func initDB(filepath string, firstInit bool) (*bbolt.DB, error) {
	db, err := bbolt.Open(filepath, dbFilePermission, &bbolt.Options{
		Timeout: DefaultPoolDBTimeout,
	})
	if err == bbolt.ErrTimeout {
		return nil, fmt.Errorf("error while trying to open %s: timed "+
			"out after %v when trying to obtain exclusive lock - "+
			"make sure no other pool daemon process (standalone "+
			"or embedded in lightning-terminal) is running",
			filepath, DefaultPoolDBTimeout)
	}
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		if firstInit {
			metadataBucket, err := tx.CreateBucketIfNotExists(
				metadataBucketKey,
			)
			if err != nil {
				return err
			}
			err = setDBVersion(metadataBucket, latestDBVersion)
			if err != nil {
				return err
			}
			if err := storeRandomLockID(metadataBucket); err != nil {
				return err
			}
		}

		_, err = tx.CreateBucketIfNotExists(accountBucketKey)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(ordersBucketKey)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(sidecarsBucketKey)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(batchBucketKey)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(eventBucketKey)
		if err != nil {
			return err
		}
		snapshotBucket, err := tx.CreateBucketIfNotExists(
			batchSnapshotBucketKey,
		)
		if err != nil {
			return err
		}

		_, err = snapshotBucket.CreateBucketIfNotExists(
			batchSnapshotSeqBucketKey,
		)
		if err != nil {
			return err
		}

		_, err = snapshotBucket.CreateBucketIfNotExists(
			batchSnapshotBatchIDIndexBucketKey,
		)
		return err
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
