package service

import (
	"context"
	"cs739-kv-store/repository"
	"database/sql"
	"log"
	"sync"
)

type PutService struct {
	memoryRepo *repository.MemoryRepo
	rdsRepo    *repository.RDSRepo
}

func NewPutService(cmap *sync.Map, db *sql.DB) *PutService {
	return &PutService{
		memoryRepo: repository.NewMemoryRepo(cmap),
		rdsRepo:    repository.NewRDSRepo(db),
	}
}

func (s *PutService) Put(ctx context.Context, key string, value string) (string, bool, error) {
	if s.memoryRepo == nil {
		return "", false, nil
	}

	oldValue, found, err := s.memoryRepo.Get(key)
	if err != nil {
		log.Printf("Error getting key: %s from memory: %v\n", key, err)
	}

	go func() {
		if err := s.memoryRepo.Put(key, value); err != nil {
			log.Printf("Error putting key: %s with value: %s in memory: %v\n", key, value, err)
		}
	}()

	if err := s.rdsRepo.Put(key, value); err != nil {
		log.Printf("Error putting key: %s with value: %s in RDS: %v\n", key, value, err)
		return "", false, err
	}

	return oldValue, found, nil
}
