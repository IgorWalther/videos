package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := new(sync.WaitGroup)
	getCfg := sync.OnceValue(getConfig)

	for range 100 {
		wg.Go(func() {
			fmt.Println(getCfg().value)
		})
	}

	wg.Wait()
}

type config struct {
	value int
}

var called atomic.Bool

func getConfig() *config {
	if called.Load() {
		panic("invariant error")
	}

	called.Store(true)
	return &config{
		value: 42,
	}
}
