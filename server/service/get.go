package service

import (
	"context"
	"cs739-kv-store/repository"
	"database/sql"
	"errors"
	"log"
)

var (
	ErrMemoryRepoNotInitialized = errors.New("memory repository is not initialized")
	ErrRDSRepoNotInitialized    = errors.New("RDS repository is not initialized")
)

type GetService struct {
	memoryRepo *repository.MemoryRepo
	rdsRepo    *repository.RDSRepo
}

func NewGetService(db *sql.DB) *GetService {
	return &GetService{
		memoryRepo: repository.NewMemoryRepo(),
		rdsRepo:    repository.NewRDSRepo(db),
	}
}

func (s *GetService) GetByKey(ctx context.Context, key string) (string, bool, error) {
	if s.memoryRepo == nil {
		return "", false, ErrMemoryRepoNotInitialized
	}

	value, found, err := s.memoryRepo.Get(key)
	if err != nil {
		return "", false, err
	}

	if found {
		// Key found in memoryRepo
		return value, true, nil
	}

	// Key not found in memoryRepo, fetch from rdsRepo
	value, found, err = s.GetByKeyFromRDS(ctx, key)
	if err != nil {
		return "", false, err
	}

	if !found {
		// Key does not exist in rdsRepo
		return "", false, nil
	}

	// Update memoryRepo with the value from rdsRepo
	_, _, err = s.memoryRepo.Put(key, value)
	if err != nil {
		log.Printf("Error putting key: %s with value: %s in memory: %v\n", key, value, err)
	}

	return value, true, nil
}

func (s *GetService) GetByKeyFromRDS(ctx context.Context, key string) (string, bool, error) {
	if s.rdsRepo == nil {
		return "", false, ErrRDSRepoNotInitialized
	}

	value, found, err := s.rdsRepo.Get(key)
	if err != nil {
		return "", false, err
	}

	return value, found, nil
}
