package repository

import (
	"sync"
)

type MemoryRepo struct {
	cmap *sync.Map
}

func NewMemoryRepo(cmap *sync.Map) *MemoryRepo {
	return &MemoryRepo{
		cmap: cmap,
	}
}

func (m *MemoryRepo) Put(key, value string) error {
	m.cmap.Store(key, value)
	return nil
}

func (m *MemoryRepo) Get(key string) (string, bool, error) {
	value, ok := m.cmap.Load(key)
	if !ok {
		return "", false, nil
	}
	return value.(string), true, nil
}
