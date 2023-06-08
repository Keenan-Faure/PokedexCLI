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
	Mux     *sync.Mutex
}

//interval is a custom value that can be set manually
// var duration time.Duration = 10000000000
// fmt.Println(duration.Seconds())
func (c Cache) NewCache(interval time.Duration, val []byte) CacheEntry {
	c.Reaploop(interval)
	return CacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func IsOldCache(cachedTime time.Time, interval time.Duration) bool {
	currentTime := time.Now()
	difference := currentTime.Sub(cachedTime)
	return difference.Seconds() > interval.Seconds()
}

func (c Cache) Add(key string, val []byte) {
	c.Mux.Lock()
	c.Mapper[key] = CacheEntry {
		createdAt: time.Now(),
		val: val,
	}
	c.Mux.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mux.Lock()
	_, ok := c.Mapper[key]
	if(ok) {
		return c.Mapper[key].val, true
	}
	c.Mux.Unlock()
	return nil, false
}

func (c Cache) Reaploop(interval time.Duration) {
	c.Mux.Lock()
	for key := range c.Mapper {
		if(IsOldCache(c.Mapper[key].createdAt, interval)) {
			delete(c.Mapper, key)
		}
		continue
	}
	c.Mux.Unlock()
}

//add the Add func
//add the get func
	//should only add the []byte from the resp.Body
//add reaploop func
	//remove older entries that are older than the interval
