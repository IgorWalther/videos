package main

import (
	"math/rand/v2"
	"sync"
)

// We can use just global lock, Microsoft

type Client struct {
	mx     *sync.Mutex
	id     int64
	amount int64
}

func bankTransfer(left, right Client) {
	// G0: 1 -> 2 (100)
	// G1: 2 -> 1 (100)

	firstMx, secondMx := left.mx, right.mx

	firstMx.Lock()
	defer firstMx.Unlock()

	// G0: lock(1)
	// G1: lock(2)

	secondMx.Lock() // G0: lock(2), G1: lock(1)
	defer secondMx.Unlock()

	right.amount += 100
	left.amount -= 100
}

func main() {
	wg := sync.WaitGroup{}

	first := Client{
		mx:     new(sync.Mutex),
		id:     1,
		amount: 1000,
	}

	second := Client{
		mx:     new(sync.Mutex),
		id:     2,
		amount: 5000,
	}

	for range 1000 {
		wg.Go(func() {
			if rand.N[int](100)%2 == 0 {
				bankTransfer(first, second)
				return
			}

			bankTransfer(second, first)
		})
	}

	wg.Wait()
}
