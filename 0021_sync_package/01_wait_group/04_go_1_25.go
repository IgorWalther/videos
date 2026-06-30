package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	var v atomic.Int64

	wg := new(WaitGroupGo1_25)
	for range 100 {
		wg.Go(func() {
			v.Add(1)
		})
	}

	wg.Wait()
	fmt.Println(v.Load())
}

type WaitGroupGo1_25 struct {
	state atomic.Uint64
}

func (w *WaitGroupGo1_25) Add(delta int) {
	c := w.state.Add(uint64(delta))

	if c < 0 {
		panic("negative state in wait group")
	}
}

func (w *WaitGroupGo1_25) Done() {
	w.Add(-1)
}

func (w *WaitGroupGo1_25) Wait() {
	for w.state.Load() != 0 {
		runtime.Gosched()
	}
}

func (w *WaitGroupGo1_25) Go(f func()) {
	w.Add(1)

	go func() {
		defer w.Done()
		f()
	}()
}
