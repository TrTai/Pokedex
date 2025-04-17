package pokecache

import (
	"time"
)

type Cache struct {
	createdAt time.Time
	val       []byte
}

func NewCache(d time.Time) *Cache {
	interval := time.Duration(d)
	newCache := Cache{
		createdAt: time.Now(),
	}
	newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) Add(k string, v []byte) {
	c.val = v
}

func (c *Cache) Get(k string) ([]byte, bool) {
	if c.val == nil {
		return nil, false
	}
	return c.val, true
}

func (c *Cache) reapLoop(d time.Duration) {
}
