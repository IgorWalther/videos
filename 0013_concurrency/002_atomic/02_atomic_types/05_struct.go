package main

import (
	"fmt"
	"sync/atomic"
)

type T struct {
	x int
}

func main() {
	var pT atomic.Pointer[T]
	ta := T{
		x: 1,
	}

	tb := T{
		x: 2,
	}

	// store
	pT.Store(&ta)
	fmt.Println(pT.Load()) // &{1}

	// load
	pa1 := pT.Load()
	fmt.Println(pa1 == &ta) // true

	// swap
	pa2 := pT.Swap(&tb)

	fmt.Println(pa2 == &ta) // true
	fmt.Println(pT.Load())  // &{2}

	// compare and swap
	b := pT.CompareAndSwap(&ta, &tb)
	fmt.Println(b) // false

	b = pT.CompareAndSwap(&tb, &ta)
	fmt.Println(b) // true
}
