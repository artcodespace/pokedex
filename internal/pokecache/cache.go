package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex sync.Mutex
}

func (c Cache) Add(key string, value []byte) {
	entry := cacheEntry {
		createdAt: time.Now()
		val: value
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	val, ok := c.entries[key]

	if ok == true {
		return val.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop(interval time.Duration)  {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	val, ok := c.entries[key]

	if ok == true {
		return val.val, true
	} else {
		return nil, false
	}
}


func NewCache(interval time.Duration) Cache {
	cache := Cache {
		entries: make(map[string]cacheEntry),
	}

	cache.reapLoop(interval)

	return cache
}

