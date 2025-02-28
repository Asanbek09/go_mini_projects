package caching

import "sync"

type Cache[K comparable, V any] struct {
	mu sync.Mutex
	data map[K]V
}

func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

func (c *Cache[K, V]) Read(key K) (V, bool) {
	v, found := c.data[key]
	return v, found
}

func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.data[key] = value

	return nil
}

func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}