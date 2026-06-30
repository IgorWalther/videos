package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

// Java - CountDownLatch
// C# - CountdownEvent
// C++ - std::latch
// Kotlin - Java CountDownLatch

func main() {
	var (
		v       atomic.Int64
		counter atomic.Int64
	)

	for range 100 {
		counter.Add(1)
		go func() {
			defer counter.Add(-1)

			v.Add(1)
		}()
	}

	for counter.Load() != 0 {
		runtime.Gosched()
	}

	fmt.Println(v.Load())
}
