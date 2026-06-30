package main

import (
	"sync"
	"time"
)

// panic: sync: WaitGroup is reused before previous Wait has returned

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	for range 100_000 {
		go func() {
			wg.Wait()
		}()
	}

	time.Sleep(time.Second)
	wg.Done()

	for range 100 {
		wg.Add(1)
	}
}
