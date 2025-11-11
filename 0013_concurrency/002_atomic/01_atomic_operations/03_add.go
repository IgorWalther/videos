package main

import (
	"fmt"
	"time"
)

func main() {
	var n int32

	// see race detector
	for range 1000 {
		go func() {
			//atomic.AddInt32(&n, 1)
			n++
		}()
	}

	time.Sleep(time.Minute) // need wg

	//fmt.Println(atomic.LoadInt32(&n)) // 1000
	fmt.Println(n)
}
