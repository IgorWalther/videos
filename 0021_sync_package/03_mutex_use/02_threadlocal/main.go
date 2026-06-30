package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	for range 100 {
		wg.Go(func() { // core 1 // core 2
			value := 0 // threadlocal

			for range 100 {
				mutex.Lock() // is it necessary?
				value++
				mutex.Unlock()
			}

			fmt.Println(value)
		})
	}

	wg.Wait()
}
