package pokecache

import (
	"sync"
	"time"
)

const ttl = time.Second * 5

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry),
	}

	ticker := time.NewTicker(interval)
	go cache.reapLoop(ticker)

	return cache
}

func (c *Cache) reapLoop(ticker *time.Ticker) {
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) >= ttl {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	res, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return res.val, ok
}
