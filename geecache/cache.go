package geecache

import (
	"go_simple_examples/geecache/lru"
	"sync"
)

type cache struct {
	cache *lru.Cache
	mutex sync.Mutex
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cache == nil {
		c.cache = lru.New(c.cacheBytes,nil)
	}
	c.cache.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cache == nil {
		return
	}
	if val, ok := c.cache.Get(key); ok {
		return val.(ByteView),ok
	}
	return
}
