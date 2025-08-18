package mutexbench

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var _ sync.Locker = (*spinLockOptimized)(nil)

func NewSpinLockOptimized() *spinLockOptimized {
	return &spinLockOptimized{}
}

type spinLockOptimized struct {
	state atomic.Int64
}

const (
	unlocked = 0
	locked   = 1
)

func (s *spinLockOptimized) Lock() {
	for {
		if s.state.CompareAndSwap(unlocked, locked) {
			return
		}

		runtime.Gosched() // time.Sleep
	}
}

func (s *spinLockOptimized) Unlock() {
	s.state.Store(unlocked)
}
