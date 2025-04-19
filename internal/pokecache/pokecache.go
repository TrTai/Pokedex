package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(d time.Duration) *Cache {
	newCache := Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(d)
	return &newCache

}

func (c *Cache) Add(k string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[k] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cacheEntries[k]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.reap(d)
		}
	}
}

func (c *Cache) reap(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cacheEntries {
		if time.Since(v.createdAt) > d {
			delete(c.cacheEntries, k)
		}
	}
}
