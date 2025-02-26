package main

type Cache[K comparable, V any] struct {
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