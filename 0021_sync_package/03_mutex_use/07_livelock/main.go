package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex1 sync.Mutex
var mutex2 sync.Mutex

func goroutine1() {
	mutex1.Lock()

	for range 10 {
		runtime.Gosched()
	}

	for !mutex2.TryLock() {
	}

	mutex2.Unlock()
	mutex1.Unlock()

	fmt.Println("G1 done")
}

func goroutine2() {
	mutex2.Lock()

	for range 10 {
		runtime.Gosched()
	}

	for !mutex1.TryLock() {
	}

	mutex1.Unlock()
	mutex2.Unlock()

	fmt.Println("G2 done")
}

func main() {
	wg := new(sync.WaitGroup)

	for range 100 {

		go func() {
			wg.Go(func() {
				goroutine1()
			})
		}()

		go func() {
			wg.Go(func() {
				goroutine2()
			})
		}()

		wg.Wait()
	}
}
