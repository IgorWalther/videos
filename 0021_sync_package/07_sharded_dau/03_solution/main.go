package dau

import (
	"hash/fnv"
	"iter"
	"sync"
	"sync/atomic"
)

type Counter interface {
	AddUser(id string) bool
	Count() uint64
	Reset()
	ForEach() iter.Seq[uint64]
}

type counterImpl struct {
	shards []shard
	total  atomic.Uint64
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

// map
// 1 2 3 4 5 6 7 8 9 10

// 1: 1 2 3 4 10
// 2: 5 6 7 8 9

// AddUser records a user by name. Returns true if the user was new today
func (c *counterImpl) AddUser(id string) bool {
	shardID := c.shardIndexFromUint(hashString64(id))
	sh := &c.shards[shardID]

	ok := withLockValues[bool](sh.mx.RLocker(), func() bool {
		_, ok := sh.set[id]
		return ok
	})

	//sh.mx.RLock()
	//_, ok := sh.set[id]
	//sh.mx.RUnlock()

	if ok {
		return false
	}

	result := withLockValues[bool](&sh.mx, func() bool {
		if _, ok := sh.set[id]; ok {
			sh.mx.Unlock()
			return false
		}

		sh.set[id] = struct{}{}
		return true
	})

	//sh.mx.Lock()
	//if _, ok := sh.set[id]; ok {
	//	sh.mx.Unlock()
	//	return false
	//}
	//
	//sh.set[id] = struct{}{}
	//sh.mx.Unlock()

	c.total.Add(1)
	return result
}

// Count returns the number of unique users seen today.
func (c *counterImpl) Count() uint64 {
	return c.total.Load()
}

// Reset clears all state (start a fresh day).
func (c *counterImpl) Reset() {
	for i := range c.shards {
		sh := &c.shards[i]

		withLock(&sh.mx, func() {
			clear(sh.set)
		})
	}

	c.total.Store(0)
}

// ForEach returns DAU iterator.
func (c *counterImpl) ForEach() iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := range c.shards {
			sh := &c.shards[i]

			withLock(sh.mx.RLocker(), func() {
				for id := range sh.set {
					if !yield(id) {
						return
					}
				}
			})
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

func withLock(l sync.Locker, action func()) {
	if action == nil {
		return
	}

	l.Lock()
	defer l.Unlock()

	action()
}

func withLockValues[T any](l sync.Locker, action func() T) T {
	l.Lock()
	defer l.Unlock()

	return action()
}
