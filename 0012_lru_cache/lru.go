package main

import (
	"container/list"
	"iter"
)

// https://leetcode.com/problems/lru-cache/

type LRUCache[K comparable, V any] interface {
	// Put O(1)
	Put(key K, value V)

	// Get O(1)
	Get(key K) V

	// Size O(1)
	Size() int

	All() iter.Seq2[K, V]
}

type node[K comparable, V any] struct {
	key   K
	value V
}

type cacheImpl[K comparable, V any] struct {
	linkedList   *list.List
	keyToElement map[K]*list.Element
	capacity     int
	defaultValue V
}

func New[K comparable, V any](
	capacity int,
	defaultValue V,
) *cacheImpl[K, V] {
	return &cacheImpl[K, V]{
		linkedList:   list.New(),
		keyToElement: make(map[K]*list.Element, capacity),
		capacity:     capacity,
		defaultValue: defaultValue,
	}
}

func (c *cacheImpl[K, V]) extractLatest() {
	del := c.linkedList.Back()
	c.linkedList.Remove(del)

	delete(c.keyToElement, c.getNodeFromElement(del).key)
}

// Put O(1)
func (c *cacheImpl[K, V]) Put(key K, value V) {
	if link, ok := c.keyToElement[key]; ok {
		n := c.getNodeFromElement(link)
		n.value = value

		c.linkedList.MoveToFront(link)
		return
	}

	if c.Size() == c.capacity {
		c.extractLatest()
	}

	n := &node[K, V]{
		key:   key,
		value: value,
	}

	c.keyToElement[key] = c.linkedList.PushFront(n)
}

func (c *cacheImpl[K, V]) getNodeFromElement(element *list.Element) *node[K, V] {
	switch v := element.Value.(type) {
	case *node[K, V]:
		return v
	default:
		panic("unreachable")
	}
}

// Get O(1)
func (c *cacheImpl[K, V]) Get(key K) V {
	if link, ok := c.keyToElement[key]; ok {
		c.linkedList.MoveToFront(link)
		return c.getNodeFromElement(link).value
	}

	return c.defaultValue
}

// Size O(1)
func (c *cacheImpl[K, V]) Size() int {
	return c.linkedList.Len()
}

func (c *cacheImpl[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		cur := c.linkedList.Front()

		for range c.Size() {
			n := c.getNodeFromElement(cur)

			if !yield(n.key, n.value) {
				return
			}

			cur = cur.Next()
		}
	}
}
	