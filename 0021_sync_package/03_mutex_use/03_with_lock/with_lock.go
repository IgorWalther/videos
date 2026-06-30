package main

import (
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	mx := new(sync.RWMutex)

	mx.Lock()
	mx.Unlock()

	mx.RLock()
	mx.RUnlock()

	repo := NewRepository()

	withLock(mx, func() {
		repo.WriteOperation()
	})

	withLock(mx.RLocker(), func() {
		repo.ReadOperation()
	})
}

// default in Kotlin
func withLock(l sync.Locker, action func()) {
	if action == nil {
		return
	}

	l.Lock()
	defer l.Unlock()

	action()
}

func NewRepository() *repository {
	return &repository{
		m: make(map[int]string),
	}
}

type repository struct {
	m map[int]string
}

//go:noinline
func (r *repository) WriteOperation() {
	r.m[rand.N[int](10e9)] = "hello"
	time.Sleep(time.Millisecond)
}

//go:noinline
func (r *repository) ReadOperation() {
	a := r.m[rand.N[int](10e9)]
	a += "///"

	time.Sleep(time.Millisecond)
}
