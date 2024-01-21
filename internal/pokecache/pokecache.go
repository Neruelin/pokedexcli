package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	store map[string]CacheEntry
	mu sync.Mutex
}

type CacheEntry struct {
	value []byte
	createdAt time.Time
}

func (pc *PokeCache) Get(url string) ([]byte, bool) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	if entry, ok := pc.store[url]; ok {
		// refreshes time to live on cache hit
		entry.createdAt = time.Now()
		return entry.value, true
	} else {
		return []byte{}, false
	}
	
}

func (pc *PokeCache) Set(url string, response []byte) {
	pc.mu.Lock()
	pc.mu.Unlock()
	pc.store[url] = CacheEntry{value: response, createdAt: time.Now()}
	return
}

func (pc *PokeCache) evictionLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		pc.eviction(time.Now().UTC().Add(-interval))
	}
}

func (pc *PokeCache) eviction(keepThreshold time.Time) {
	pc.mu.Lock()
	pc.mu.Unlock()

	for k, v := range pc.store {
		if v.createdAt.Before(keepThreshold) {
			delete(pc.store, k)
		}
	}
}

func New() PokeCache {
	pc := PokeCache{store: map[string]CacheEntry{}, mu: sync.Mutex{}}
	go pc.evictionLoop(60 * time.Second)
	return pc
}