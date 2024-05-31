package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val      []byte
	createAt time.Time
}

func NewCacheHelper(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:      val,
		createAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	// this will create a timer that will run the reap function at the duration interval:w
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
