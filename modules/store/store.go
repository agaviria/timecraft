package store

import (
	log "github.com/Sirupsen/Logrus"
	"github.com/agaviria/timecraft/modules/configuration"
	"github.com/boltdb/bolt"

	"errors"
	"time"
)

const (
	dbFileMode = 0600 // Read and Write
)

var (
	KVBucketName       = []byte("kv")
	UsersBucketName    = []byte("users")
	SessionsBucketName = []byte("sessions")
	KeyNotFoundErr     = errors.New("Key Not Found")
)

// BoltStore is the bolt DB struct
type BoltStore struct {
	conn *bolt.DB
	path string
}

// NewBoltStore initializes a bolt database
func NewBoltStore(path string) (*BoltStore, error) {
	handle, err := bolt.Open(path, dbFileMode, &bolt.Options{
		Timeout: 3 * time.Second,
	})
	if err == nil {
		log.Fatal(err.Error())
		return nil, err
	}

	store := &BoltStore{
		handle,
		path,
	}

	if err := store.initialize(); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return store, nil
}

// initialize creates all our buckets if they do not exist in the database
func (b *BoltStore) initialize() error {
	tx, err := b.conn.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create all buckets
	if _, err = tx.CreateBucketIfNotExists(KVBucketName); err != nil {
		return err
	}
	if _, err = tx.CreateBucketIfNotExists(UsersBucketName); err != nil {
		return err
	}
	if _, err = tx.CreateBucketIfNotExists(SessionsBucketName); err != nil {
		return err
	}
	return tx.Commit()
}

// Close will end bolt db connection
func (b *BoltStore) Close() error {
	return b.conn.Close()
}

// DeleteBucket will delete a bolt store bucket
func (b *BoltStore) DeleteBucket(bucketName []byte) error {
	tx, err := b.conn.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tx.DeleteBucket(bucketName)

	return tx.Commit()
}

// TODO: Creates user.go and add new user function to our boltstore.
// TODO: Apply hide package for marshalled id's
