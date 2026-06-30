package main

import (
	"fmt"
	"sync"

	"github.com/petermattis/goid"
)

// TODO: Go 1.25
func main() {
	m := make(map[int64]struct{})
	mx := new(sync.Mutex)

	wg := new(sync.WaitGroup)
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mx.Lock()
			defer mx.Unlock()

			id := goid.Get()
			m[id] = struct{}{}

			fmt.Println(id)
		}()
	}

	wg.Wait()

	if len(m) != 1000 {
		panic("invariant error")
	}
}
