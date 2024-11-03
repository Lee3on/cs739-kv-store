package service

import (
	"bytes"
	"cs739-kv-store/consts"
	"cs739-kv-store/raft"
	"cs739-kv-store/repository"
	"database/sql"
	"encoding/gob"
	"errors"
	"log"
	"strings"
	"sync"

	"go.etcd.io/etcd/raft/v3/raftpb"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
)

// a key-value store backed by raft
type Kvstore struct {
	proposeC chan<- string // channel for proposing updates
	mu       sync.RWMutex
	//kvStore     map[string]string // current committed key-value pairs
	memoryRepo  *repository.MemoryRepo
	rdsRepo     *repository.RDSRepo
	snapshotter *snap.Snapshotter
}

type kv struct {
	Key string
	Val string
}

func NewKVStore(snapshotter *snap.Snapshotter, proposeC chan<- string, commitC <-chan *raft.Commit, errorC <-chan error, db *sql.DB) *Kvstore {
	s := &Kvstore{
		proposeC: proposeC,
		//kvStore:     make(map[string]string),
		memoryRepo:  repository.NewMemoryRepo(consts.KVStoreCapacity, consts.KVStoreEvictionTTL),
		rdsRepo:     repository.NewRDSRepo(db),
		snapshotter: snapshotter,
	}
	snapshot, err := s.loadSnapshot()
	if err != nil {
		log.Panic(err)
	}
	if snapshot != nil {
		log.Printf("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
		if err := s.recoverFromSnapshot(snapshot.Data); err != nil {
			log.Panic(err)
		}
	}
	// read commits from raft into kvStore map until error
	go s.readCommits(commitC, errorC)
	return s
}

func (s *Kvstore) Get(key string) (string, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return NewGetService(s.memoryRepo, s.rdsRepo).GetByKey(key)
}

func (s *Kvstore) Put(k string, v string) (string, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	oldValue, found, err := NewGetService(s.memoryRepo, s.rdsRepo).GetByKey(k)
	s.Propose(k, v)
	return oldValue, found, err
}

func (s *Kvstore) Propose(k string, v string) {
	var buf strings.Builder
	if err := gob.NewEncoder(&buf).Encode(kv{k, v}); err != nil {
		log.Fatal(err)
	}
	s.proposeC <- buf.String()
}

func (s *Kvstore) readCommits(commitC <-chan *raft.Commit, errorC <-chan error) {
	for commit := range commitC {
		if commit == nil {
			// signaled to load snapshot
			snapshot, err := s.loadSnapshot()
			if err != nil {
				log.Panic(err)
			}
			if snapshot != nil {
				log.Printf("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
				if err := s.recoverFromSnapshot(snapshot.Data); err != nil {
					log.Panic(err)
				}
			}
			continue
		}

		for _, data := range commit.Data {
			var dataKv kv
			dec := gob.NewDecoder(bytes.NewBufferString(data))
			if err := dec.Decode(&dataKv); err != nil {
				log.Fatalf("raftexample: could not decode message (%v)", err)
			}
			s.mu.Lock()
			if err := NewPutService(s.memoryRepo, s.rdsRepo).Put(dataKv.Key, dataKv.Val); err != nil {
				log.Fatalf("Error putting key: %s with value: %s in memory: %v\n", dataKv.Key, dataKv.Val, err)
			}
			s.mu.Unlock()
		}
		close(commit.ApplyDoneC)
	}
	if err, ok := <-errorC; ok {
		log.Fatal(err)
	}
}

func (s *Kvstore) GetSnapshot() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.rdsRepo.Serialize()
}

func (s *Kvstore) loadSnapshot() (*raftpb.Snapshot, error) {
	snapshot, err := s.snapshotter.Load()
	if errors.Is(err, snap.ErrNoSnapshot) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return snapshot, nil
}

func (s *Kvstore) recoverFromSnapshot(snapshot []byte) error {
	return s.rdsRepo.Deserialize(snapshot)
}

func (s *Kvstore) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key, elem := range s.memoryRepo.GetCache() {
		entry := elem.Value.(*repository.CacheEntry)
		value := entry.Value
		err := s.rdsRepo.Put(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
