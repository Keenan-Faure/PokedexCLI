package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Mapper map[string]CacheEntry
	Mux    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	Cache := Cache{
		Mapper: make(map[string]CacheEntry),
		Mux:    &sync.Mutex{},
	}
	go Cache.Reaploop(interval)
	return Cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	c.Mapper[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	_, ok := c.Mapper[key]
	if ok {
		return c.Mapper[key].val, true
	}
	return nil, false
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	c.Mux.Lock()
	defer c.Mux.Unlock()
	for key, value := range c.Mapper {
		if value.createdAt.Before(timeAgo) {
			delete(c.Mapper, key)
		}
	}
}

func (c *Cache) Reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}
