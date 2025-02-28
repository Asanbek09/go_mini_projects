package caching

import (
	"sync"
	"time"
)

type entryWithTimeout[V any] struct {
	value V
	expires time.Time
}

type Cache[K comparable, V any] struct {
	ttl time.Duration

	mu sync.RWMutex
	data map[K]entryWithTimeout[V]
}

func New[K comparable, V any](ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl: ttl,
		data: make(map[K]entryWithTimeout[V]),
	}
}

func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	v, found := c.data[key]
	return v, found
}

func (c *Cache[K, V]) Upsert(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.data[key] = entryWithTimeout[V]{
		value: value,
		expires: time.Now().Add(c.ttl),
	}
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}