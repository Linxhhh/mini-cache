package mini_cache

import (
	"github.com/Linxhhh/mini-cache/lru"
	"sync"
)

// cache 封装了 lru 缓存，实现了缓存数据的并发访问

type Cache struct {
	mu        sync.RWMutex
	lru       *lru.Cache
	cacheByte int64
}

func NewCache(cacheByte int64) *Cache {
	return &Cache{
		lru: nil,
		cacheByte: cacheByte,
	}
}

func (c *Cache) Add(key string, value Byteview) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheByte)
	}
	c.lru.Add(key, value)
}

func (c *Cache) Get(key string) (Byteview, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.lru == nil {
		return Byteview{}, false
	}
	v, ok := c.lru.Get(key)
	if ok {
		return v.(Byteview), ok
	}
	return Byteview{}, ok
}