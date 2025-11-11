package main

import (
	"sync/atomic"
	"time"
)

// example1
func loadConfig() map[string]string {
	return make(map[string]string)
}

func requests() chan int {
	return make(chan int)
}

func main() {
	var cfg atomic.Value
	cfg.Store(loadConfig())

	go func() {
		for {
			time.Sleep(10 * time.Second)
			cfg.Store(loadConfig())
		}
	}()

	for r := range requests() {
		c := cfg.Load().(map[string]string) // TODO: generic wrapper
		_, _ = r, c
	}
}
