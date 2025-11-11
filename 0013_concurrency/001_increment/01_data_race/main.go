package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var v int64

	for range 1000 {
		// Пока для простоты можно думать, что goroutine == os thread
		go func() {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			v++
		}()
	}

	time.Sleep(time.Second * 5) // TODO: wait group
	fmt.Println(v)              // why???
}
