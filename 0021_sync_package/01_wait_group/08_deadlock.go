package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var v atomic.Int64
	wg := new(sync.WaitGroup)

	for range 100 {
		// see wg.Go()

		wg.Add(1)
		go func() {
			// defer wg.Done() // no done!
			v.Add(1)
		}()
	}

	wg.Wait()
	// fatal error: all goroutines are asleep - deadlock!
}
