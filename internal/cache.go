package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(td time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(td)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = entry

	c.mutex.Unlock()

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		fmt.Println("Key not found in cache")
		return nil, false
	}

	return entry.val, true

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		c.mutex.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}
