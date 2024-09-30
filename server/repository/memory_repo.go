package repository

import (
	"container/list"
	"sync"
	"time"
)

// cacheEntry represents a single entry in the cache.
type cacheEntry struct {
	key        string
	value      string
	expiration time.Time
}

// MemoryRepo implements an in-memory key-value store with TTL and LRU eviction.
type MemoryRepo struct {
	mu       sync.Mutex
	capacity int
	ttl      time.Duration
	cache    map[string]*list.Element
	lruList  *list.List
}

// NewMemoryRepo creates a new MemoryRepo with the given capacity and TTL.
func NewMemoryRepo(capacity int, ttl time.Duration) *MemoryRepo {
	repo := &MemoryRepo{
		capacity: capacity,
		ttl:      ttl,
		cache:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
	// Start a background goroutine to clean up expired entries.
	go repo.startEviction()
	return repo
}

// Put adds or updates a key-value pair in the cache.
func (m *MemoryRepo) Put(key, value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if elem, ok := m.cache[key]; ok {
		// Update existing entry.
		entry := elem.Value.(*cacheEntry)
		entry.value = value
		entry.expiration = time.Now().Add(m.ttl)
		m.lruList.MoveToFront(elem)
	} else {
		// Add new entry.
		entry := &cacheEntry{
			key:        key,
			value:      value,
			expiration: time.Now().Add(m.ttl),
		}
		elem := m.lruList.PushFront(entry)
		m.cache[key] = elem

		// Check capacity and evict if necessary.
		if m.lruList.Len() > m.capacity {
			m.evict()
		}
	}
	return nil
}

// Get retrieves the value for a given key.
func (m *MemoryRepo) Get(key string) (string, bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if elem, ok := m.cache[key]; ok {
		entry := elem.Value.(*cacheEntry)
		if time.Now().After(entry.expiration) {
			// Entry has expired.
			m.removeElement(elem)
			return "", false, nil
		}
		// Update LRU order.
		m.lruList.MoveToFront(elem)
		return entry.value, true, nil
	}
	return "", false, nil
}

// evict removes the least recently used item from the cache.
func (m *MemoryRepo) evict() {
	elem := m.lruList.Back()
	if elem != nil {
		m.removeElement(elem)
	}
}

// removeElement removes an element from the cache and list.
func (m *MemoryRepo) removeElement(elem *list.Element) {
	m.lruList.Remove(elem)
	entry := elem.Value.(*cacheEntry)
	delete(m.cache, entry.key)
}

// startEviction runs in the background to remove expired entries.
func (m *MemoryRepo) startEviction() {
	ticker := time.NewTicker(m.ttl / 2)
	defer ticker.Stop()
	for {
		<-ticker.C
		m.mu.Lock()
		now := time.Now()
		for elem := m.lruList.Back(); elem != nil; {
			prev := elem.Prev()
			entry := elem.Value.(*cacheEntry)
			if now.After(entry.expiration) {
				m.removeElement(elem)
			}
			elem = prev
		}
		m.mu.Unlock()
	}
}
