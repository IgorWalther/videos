package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	wg := new(sync.WaitGroup)
	once := new(sync.Once)

	for range 1_000 {
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
