package repository

import (
	"os"

	jsoniter "github.com/json-iterator/go"
	bolt "go.etcd.io/bbolt"
)

var (
	Bucket            = []byte("bucket")
	IndexKey          = []byte("index")
	StateKey          = []byte("kv_state") // Key for storing KV state
	EmptyIndex uint64 = 0
)

type Storage struct {
	Path string
	DB   *bolt.DB
}

func NewStorage(path string) (*Storage, error) {
	db, err := bolt.Open(path, os.ModePerm, nil)
	if err != nil {
		return nil, err
	}

	storage := Storage{
		Path: path,
		DB:   db,
	}
	return &storage, nil
}

func (storage *Storage) Create(index uint64) error {
	return storage.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(Bucket)
		if err != nil {
			return err
		}

		bs := bucket.Get(IndexKey)
		if len(bs) != 0 {
			return nil
		}
		bs, err = jsoniter.Marshal(index)
		if err != nil {
			return err
		}
		return bucket.Put(IndexKey, bs)
	})
}

func (storage *Storage) Put(index uint64) error {
	bs, err := jsoniter.Marshal(index)
	if err != nil {
		return err
	}
	return storage.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(Bucket)
		return bucket.Put(IndexKey, bs)
	})
}

func (storage *Storage) Get() (uint64, error) {
	var bs []byte
	err := storage.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(Bucket)
		bs = bucket.Get(IndexKey)
		return nil
	})
	if len(bs) == 0 || err != nil {
		return EmptyIndex, err
	}

	var index uint64
	err = jsoniter.Unmarshal(bs, &index)
	return index, err
}

// SaveState persists the key-value state in BoltDB
func (storage *Storage) SaveState(kv map[string]string) error {
	// Serialize the key-value map
	bs, err := jsoniter.Marshal(kv)
	if err != nil {
		return err
	}

	// Save serialized map in BoltDB
	return storage.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(Bucket)
		return bucket.Put(StateKey, bs)
	})
}

// LoadState loads the key-value state from BoltDB
func (storage *Storage) LoadState() (map[string]string, error) {
	var bs []byte
	err := storage.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(Bucket)
		bs = bucket.Get(StateKey)
		return nil
	})

	// If no state is found, return an empty map
	if err != nil || len(bs) == 0 {
		return make(map[string]string), err
	}

	// Deserialize the state
	var kv map[string]string
	err = jsoniter.Unmarshal(bs, &kv)
	return kv, err
}

func (storage *Storage) Stop() error {
	return storage.DB.Close()
}
