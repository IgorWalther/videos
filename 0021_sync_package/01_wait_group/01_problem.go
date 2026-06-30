package main

import (
	"math/rand/v2"
	"sync/atomic"
	"time"
)

func main() {
	var v atomic.Int64

	for range 1000 {
		go func() {
			v.Add(1)
			time.Sleep(rand.N[time.Duration](5000))
		}()
	}

	time.Sleep(10 * time.Second)

	//time.Sleep(time.Second * 3)
	//time.Sleep(time.Second * 5)
}
