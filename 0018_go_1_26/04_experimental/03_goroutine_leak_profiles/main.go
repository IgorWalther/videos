//go:build goexperiment.goroutineleakprofile

package main

import (
	"os"
	"runtime/pprof"
	"time"
)

// We aim to enable goroutine leak profiles by default in Go 1.27.

func main() {
	prof := pprof.Lookup("goroutineleak")
	leak()
	time.Sleep(50 * time.Millisecond)
	prof.WriteTo(os.Stdout, 2)
}

func leak() <-chan int {
	out := make(chan int)

	go func() {
		out <- 42
	}()

	return out
}
