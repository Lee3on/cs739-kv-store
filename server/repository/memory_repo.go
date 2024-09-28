package repository

import "sync"

type MemoryRepo struct {
	cmap sync.Map
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		cmap: sync.Map{},
	}
}

func (m *MemoryRepo) Put(key, value string) (string, bool, error) {
	oldValue, ok := m.cmap.LoadOrStore(key, value)
	if !ok {
		return "", false, nil
	}

	return oldValue.(string), true, nil
}

func (m *MemoryRepo) Get(key string) (string, bool, error) {
	value, ok := m.cmap.Load(key)
	if !ok {
		return "", false, nil
	}
	return value.(string), true, nil
}
