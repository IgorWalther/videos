package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := new(sync.WaitGroup)
	init := sync.OnceFunc(initCfg)

	for range 100 {
		wg.Go(func() {
			init()
			fmt.Println(cfg.value)
		})
	}

	wg.Wait()
}

var called atomic.Bool

type config struct {
	value int
}

var cfg *config

func initCfg() {
	if called.Load() {
		panic("invariant error")
	}

	cfg = &config{
		value: 42,
	}

	called.Store(true)
}
