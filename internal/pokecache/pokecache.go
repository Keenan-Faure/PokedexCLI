package pokecache

import (
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration, data []byte) {

}

//add the Add func
//add the get func
	//should only add the []byte from the resp.Body
//add reaploop func
	//remove older entries that are older than the interval
