package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheEntry, ok := c.cache[key]
	return cacheEntry.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	nowMinusInterval := time.Now().UTC().Add(-interval)
	for key, value := range c.cache {
		if value.createdAt.Before(nowMinusInterval) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	val []byte
	createdAt time.Time
}