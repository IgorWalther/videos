package main

import (
	"sync"
	"sync/atomic"
)

type myOnce struct {
	done atomic.Uint32
}

func (o *myOnce) Do(f func()) {
	if o.done.CompareAndSwap(0, 1) {
		f()
	}
}

func main() {
	wg := new(sync.WaitGroup)
	once := new(myOnce)

	for range 10_000 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			once.Do(nonIdempotentClose)
		}()
	}

	wg.Wait()

	if !isCalled.Load() {
		panic("invariant error")
	}
}

var isCalled atomic.Bool

// not pure function
func nonIdempotentClose() {
	if isCalled.Load() {
		panic("invariant error")
	}

	isCalled.Store(true)
}
