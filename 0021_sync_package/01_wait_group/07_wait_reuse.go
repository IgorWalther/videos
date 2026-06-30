package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	v := atomic.Int64{}
	wg := new(sync.WaitGroup)

	for range 100 {
		wg.Go(func() {
			v.Add(1)
		})
	}

	wg.Wait()

	for range 100 {
		wg.Go(func() {
			v.Add(1)
		})
	}

	wg.Wait()
}
