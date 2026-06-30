package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := new(sync.WaitGroup)
	getCfg := sync.OnceValues(getConfig)

	for range 100 {
		wg.Go(func() {
			cfg, err := getCfg()

			fmt.Println("err: ", err, "value: ", cfg.value)
		})
	}

	wg.Wait()
}

type config struct {
	value int
}

var called atomic.Bool

func getConfig() (*config, error) {
	if called.Load() {
		panic("invariant error")
	}

	called.Store(true)
	return &config{
		value: 42,
	}, nil
}
