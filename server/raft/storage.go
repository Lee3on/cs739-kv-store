package raft

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"go.etcd.io/etcd/raft/v3/raftpb"
)

// BoltStorage implements the etcd/raft Storage interface using BoltDB.
type BoltStorage struct {
	db *bolt.DB
}

// NewBoltStorage creates a new instance of BoltStorage
func NewBoltStorage(path string) (*BoltStorage, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &BoltStorage{db: db}, nil
}

// Save persists log entries and the hard state to disk.
func (bs *BoltStorage) Save(hardState raftpb.HardState, entries []raftpb.Entry) error {
	return bs.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("raft"))
		if err != nil {
			return err
		}

		// Persist hard state
		hardStateBytes, err := hardState.Marshal()
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("hard_state"), hardStateBytes)
		if err != nil {
			return err
		}

		// Persist entries
		for _, entry := range entries {
			entryBytes, err := entry.Marshal()
			if err != nil {
				return err
			}
			err = bucket.Put(itob(entry.Index), entryBytes)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// Load retrieves the last persisted state from BoltDB.
func (bs *BoltStorage) InitialState() (raftpb.HardState, raftpb.ConfState, error) {
	var hardState raftpb.HardState
	err := bs.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("raft"))
		if bucket == nil {
			return nil // No state yet
		}

		hardStateBytes := bucket.Get([]byte("hard_state"))
		if hardStateBytes == nil {
			return nil
		}

		return hardState.Unmarshal(hardStateBytes)
	})

	return hardState, raftpb.ConfState{}, err
}

// Entries returns a slice of entries in the range [lo, hi).
func (bs *BoltStorage) Entries(lo, hi, maxSize uint64) ([]raftpb.Entry, error) {
	var entries []raftpb.Entry
	err := bs.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("raft"))
		if bucket == nil {
			return nil
		}

		for i := lo; i < hi; i++ {
			entryBytes := bucket.Get(itob(i))
			if entryBytes == nil {
				continue
			}

			var entry raftpb.Entry
			err := entry.Unmarshal(entryBytes)
			if err != nil {
				return err
			}
			entries = append(entries, entry)
		}

		return nil
	})

	return entries, err
}

// Term returns the term of the entry at index i.
func (bs *BoltStorage) Term(i uint64) (uint64, error) {
	var term uint64
	err := bs.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("raft"))
		if bucket == nil {
			return fmt.Errorf("term not found for index %d", i)
		}

		entryBytes := bucket.Get(itob(i))
		if entryBytes == nil {
			return fmt.Errorf("entry not found for index %d", i)
		}

		var entry raftpb.Entry
		err := entry.Unmarshal(entryBytes)
		if err != nil {
			return err
		}
		term = entry.Term
		return nil
	})

	return term, err
}

// LastIndex returns the index of the last entry in the log.
func (bs *BoltStorage) LastIndex() (uint64, error) {
	var lastIndex uint64
	err := bs.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("raft"))
		if bucket == nil {
			return nil // No entries
		}

		cursor := bucket.Cursor()
		key, _ := cursor.Last()
		if key != nil {
			lastIndex = btoi(key)
		}

		return nil
	})

	return lastIndex, err
}

// FirstIndex returns the index of the first entry in the log.
func (bs *BoltStorage) FirstIndex() (uint64, error) {
	var firstIndex uint64
	err := bs.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("raft"))
		if bucket == nil {
			return nil // No entries
		}

		cursor := bucket.Cursor()
		key, _ := cursor.First()
		if key != nil {
			firstIndex = btoi(key)
		}

		return nil
	})

	return firstIndex, err
}

// Snapshot returns the most recent snapshot.
func (bs *BoltStorage) Snapshot() (raftpb.Snapshot, error) {
	// For simplicity, returning an empty snapshot
	// In a real-world scenario, you would persist and return actual snapshots
	return raftpb.Snapshot{}, fmt.Errorf("not implemented")
}

// Helper function to convert uint64 to byte slice
func itob(v uint64) []byte {
	b := make([]byte, 8)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}

// btoi converts a byte slice to a uint64.
func btoi(b []byte) uint64 {
	return uint64(b[0])<<56 |
		uint64(b[1])<<48 |
		uint64(b[2])<<40 |
		uint64(b[3])<<32 |
		uint64(b[4])<<24 |
		uint64(b[5])<<16 |
		uint64(b[6])<<8 |
		uint64(b[7])
}
