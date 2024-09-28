package service

import (
	"context"
	"cs739-kv-store/repository"
	"database/sql"
	"log"
)

type PutService struct {
	memoryRepo *repository.MemoryRepo
	rdsRepo    *repository.RDSRepo
}

func NewPutService(db *sql.DB) *PutService {
	return &PutService{
		memoryRepo: repository.NewMemoryRepo(),
		rdsRepo:    repository.NewRDSRepo(db),
	}
}

func (s *PutService) Put(ctx context.Context, key string, value string) (string, bool, error) {
	if s.memoryRepo == nil {
		return "", false, nil
	}
	if err := s.rdsRepo.Put(key, value); err != nil {
		log.Printf("Error putting key: %s with value: %s in RDS: %v\n", key, value, err)
		return "", false, err
	}

	oldValue, found, err := s.memoryRepo.Put(key, value)
	if err != nil {
		log.Printf("Error putting key: %s with value: %s in memory: %v\n", key, value, err)
		return "", false, err
	}

	return oldValue, found, nil
}
