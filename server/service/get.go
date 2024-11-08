package service

import (
	"cs739-kv-store/repository"
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

func NewGetService(memoryRepo *repository.MemoryRepo, rdsRepo *repository.RDSRepo) *GetService {
	return &GetService{
		memoryRepo: memoryRepo,
		rdsRepo:    rdsRepo,
	}
}

func (s *GetService) GetByKey(key string) (string, bool, error) {
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
	log.Printf("Key: %s not found in memoryRepo, fetching from rdsRepo\n", key)
	value, found, err = s.GetByKeyFromRDS(key)
	if err != nil {
		return "", false, err
	}

	if !found {
		// Key does not exist in rdsRepo
		return "", false, nil
	}

	// Update memoryRepo with the value from rdsRepo
	if err = s.memoryRepo.Put(key, value); err != nil {
		log.Printf("Error putting key: %s with value: %s in memory: %v\n", key, value, err)
	}

	return value, true, nil
}

func (s *GetService) GetByKeyFromRDS(key string) (string, bool, error) {
	if s.rdsRepo == nil {
		return "", false, ErrRDSRepoNotInitialized
	}

	value, found, err := s.rdsRepo.Get(key)
	if err != nil {
		return "", false, err
	}

	return value, found, nil
}
