package service

import (
	"cs739-kv-store/repository"
	"log"
)

type PutService struct {
	memoryRepo *repository.MemoryRepo
	rdsRepo    *repository.RDSRepo
}

func NewPutService(memoryRepo *repository.MemoryRepo, rdsRepo *repository.RDSRepo) *PutService {
	return &PutService{
		memoryRepo: memoryRepo,
		rdsRepo:    rdsRepo,
	}
}

func (s *PutService) Put(key string, value string) error {
	if s.memoryRepo == nil {
		return nil
	}

	go func() {
		if err := s.memoryRepo.Put(key, value); err != nil {
			log.Printf("Error putting key: %s with value: %s in memory: %v\n", key, value, err)
		}
	}()

	if err := s.rdsRepo.Put(key, value); err != nil {
		log.Printf("Error putting key: %s with value: %s in RDS: %v\n", key, value, err)
		return err
	}

	return nil
}
