package main

import (
	"sync"
	"sync/atomic"
)

// C++ - std::once_flag + std::call_once
// Rust - std::sync::Once
// Kotlin - by lazy

type mySyncOnce struct {
	done atomic.Bool
}

func (m *mySyncOnce) Do(f func()) {
	if m.done.CompareAndSwap(false, true) {
		f()
	}
}

type Map struct {
	once   *mySyncOnce
	mx     *sync.Mutex
	values map[any]any
}

func NewMap() *Map {
	return &Map{
		mx:   new(sync.Mutex),
		once: new(mySyncOnce),
	}
}

func (m *Map) Add(key, value any) {
	m.init()

	m.mx.Lock()
	m.values[key] = value
	m.mx.Unlock()
}

func (m *Map) init() {
	m.once.Do(func() {
		m.values = make(map[any]any) // G1
	})
}
