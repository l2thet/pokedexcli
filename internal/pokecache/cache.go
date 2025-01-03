package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mutex:   sync.Mutex{},
	}

	go cache.reaploop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.mutex.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	entry, ok := cache.entries[key]
	cache.mutex.Unlock()

	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (cache *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		currentTime := time.Now()
		reapTime := currentTime.Add(-interval)

		cache.mutex.Lock()
		for key, entry := range cache.entries {
			if entry.createdAt.Before(reapTime) {
				delete(cache.entries, key)
			}
		}
		cache.mutex.Unlock()
	}
}
