package dau

import (
	"hash/fnv"
	"iter"
	"sync"
)

type Counter interface {
	AddUser(id string) bool
	Count() uint64
	Reset()
	ForEach() iter.Seq[uint64]
}

type counterImpl struct {
	shards []shard
}

type Set[K comparable] = map[K]struct{}

type shard struct {
	mx  sync.RWMutex
	set Set[string]
}

const defaultShardSize = 1 << 10

func NewCounter(shards int) *counterImpl {
	if shards < 1 {
		shards = 1
	}

	c := &counterImpl{
		shards: make([]shard, shards),
	}

	for i := range c.shards {
		c.shards[i].set = make(map[string]struct{}, defaultShardSize)
	}

	return c
}

// AddUser records a user by name. Returns true if the user was new today
func (c *counterImpl) AddUser(id string) bool {
	shardID := c.shardIndexFromUint(hashString64(id))
	sh := &c.shards[shardID]

	sh.mx.RLock()
	_, ok := sh.set[id]
	if ok {
		sh.mx.RUnlock()
		return false
	}
	sh.mx.RUnlock()

	// <----------------------------------------

	sh.mx.Lock()
	_, ok = sh.set[id]

	if ok {
		sh.mx.Unlock()
		return false
	}

	sh.set[id] = struct{}{}
	sh.mx.Unlock()

	return true
}

// Count returns the number of unique users seen today.
func (c *counterImpl) Count() uint64 {
	var total uint64
	for i := range c.shards {
		sh := &c.shards[i]
		sh.mx.RLock()
		total += uint64(len(c.shards))
		sh.mx.RUnlock()
	}

	return total
}

// Reset clears all state (start a fresh day).
func (c *counterImpl) Reset() {
	for i := range c.shards {
		sh := &c.shards[i]
		sh.mx.Lock()
		clear(sh.set)
		// not sh.set = make(map[string]struct{}, 1024)
		sh.mx.Unlock()
	}
}

// ForEach returns DAU iterator.
func (c *counterImpl) ForEach() iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := range c.shards {
			sh := &c.shards[i]

			sh.mx.RLock()
			for id := range sh.set {
				if !yield(id) {
					sh.mx.RUnlock()
					return
				}
			}

			sh.mx.RUnlock()
		}
	}
}

func (c *counterImpl) shardIndexFromUint(id uint64) int {
	return int(id % uint64(len(c.shards)))
}

func hashString64(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}
