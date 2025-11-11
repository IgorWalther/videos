package main

import (
	"fmt"
	"runtime"
	"time"
)

// LDADDAL
// https://developer.arm.com/documentation/111108/2025-09/Base-Instructions/LDADD--LDADDA--LDADDAL--LDADDL--Atomic-add-on-word-or-doubleword-?lang=en

//go:noescape
func SyncAdd(addr *int64, delta int64)

func main() {
	var v int64

	for range 1000 {
		// Пока для простоты можно думать, что goroutine == os thread
		go func() {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()

			SyncAdd(&v, 1) // v++
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println(v) // TODO: wait group
}
