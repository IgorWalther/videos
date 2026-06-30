package dau

import (
	"iter"
	"sync"
)

// TODO: uuid

type Counter interface {
	AddUser(id string) bool
	Count() uint64
	Reset()
	ForEach() iter.Seq[string]
}

var _ Counter = (*counterImpl)(nil)

type counterImpl struct {
	data map[string]struct{} // Go 1.24+ (Go 1.27)
	mx   sync.Mutex
}

func NewCounter() *counterImpl {
	return &counterImpl{
		data: make(map[string]struct{}),
	}
}

// AddUser records a user by name. Returns true if the user was new today
func (c *counterImpl) AddUser(id string) bool {
	c.mx.Lock()
	defer c.mx.Unlock()

	_, ok := c.data[id]

	if ok {
		return true
	}

	c.data[id] = struct{}{}
	return false
}

// Count returns the number of unique users seen today.
func (c *counterImpl) Count() uint64 {
	c.mx.Lock()
	defer c.mx.Unlock()

	return uint64(len(c.data))
}

// Reset clears all state (start a fresh day).
func (c *counterImpl) Reset() {
	c.mx.Lock()
	defer c.mx.Unlock()
	clear(c.data)
}

// ForEach returns DAU iterator.
func (c *counterImpl) ForEach() iter.Seq[string] {
	return func(yield func(string) bool) {
		c.mx.Lock()
		defer c.mx.Unlock()

		for k := range c.data {
			if yield(k) {
				return
			}
		}
	}
}
