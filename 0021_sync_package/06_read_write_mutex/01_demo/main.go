package main

import "sync"

type Repository[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (c *Repository[K, V]) Load(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	return value, found
}

func (c *Repository[K, V]) Store(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value
}
