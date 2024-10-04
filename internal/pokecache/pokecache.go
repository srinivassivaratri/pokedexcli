package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

// our package users will use this function to create a new cache. 
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}


// Add() method adds a new key-value pair(i.e. new entry to the cache).
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val: value, 
		createdAt: time.Now().UTC(),
	}
}

// Get() method retrieves a value from the cache based on the provided key.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

// In layman terms, reapLoop() method is a function that runs in a loop, checking for expired cache entries and removing them.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}


// In layman terms, reap() method is a function that checks for expired cache entries and deletes them.
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}





