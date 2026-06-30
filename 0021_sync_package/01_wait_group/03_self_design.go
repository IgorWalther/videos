package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	var v atomic.Int64

	wg := new(WaitGroup)
	for range 100 {
		wg.Add(1) // slower, but readable
		go func() {
			defer wg.Done()

			v.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println(v.Load())
}

type WaitGroup struct {
	state atomic.Uint64
}

func (w *WaitGroup) Add(delta int) {
	c := w.state.Add(uint64(delta))

	if c < 0 {
		panic("negative state in wait group")
	}
}

func (w *WaitGroup) Done() {
	w.Add(-1)
}

func (w *WaitGroup) Wait() {
	for w.state.Load() != 0 {
		runtime.Gosched()
	}
}
