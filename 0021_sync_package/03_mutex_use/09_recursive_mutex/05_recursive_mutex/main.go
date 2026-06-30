package main

import (
	"errors"
	"runtime"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// reentrant

var ErrUnlockFromAnotherGoroutine = errors.New("unlock from non-owner goroutine")

func New() *Mutex {
	s := atomic.Int64{}
	s.Store(unlocked)

	return &Mutex{
		ownerID: &s,
	}
}

const unlocked = -1

type Mutex struct {
	ownerID    *atomic.Int64
	ownerCalls atomic.Int64
}

func (r *Mutex) Lock() {
	goroutineID := goid.Get()

	if r.ownerID.Load() == goroutineID {
		r.ownerCalls.Add(1)
		return
	}

	for !r.ownerID.CompareAndSwap(unlocked, goroutineID) {
		runtime.Gosched()
	}

	r.ownerCalls.Store(1)
}

func (r *Mutex) Unlock() {
	goroutineID := goid.Get()

	if r.ownerID.Load() != goroutineID {
		panic(ErrUnlockFromAnotherGoroutine)
	}

	if r.ownerCalls.Add(-1) == 0 {
		r.ownerID.Store(unlocked)
	}
}
