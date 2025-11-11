package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type T struct {
		a int
		b int
		c int
	}

	x := T{
		a: 1,
		b: 2,
		c: 3,
	}

	y := T{
		a: 4,
		b: 5,
		c: 6,
	}

	z := T{
		a: 7,
		b: 8,
		c: 9,
	}

	var v atomic.Value

	fmt.Println(v)   // {{1 2 3}}
	old := v.Swap(y) // for example - cfg

	fmt.Println(v)                    // {{4 5 6}}
	fmt.Println(old.(T))              // {1 2 3}
	swapped := v.CompareAndSwap(x, z) // Go 1.17
	fmt.Println(swapped, v)           // false {{4 5 6}}
	swapped = v.CompareAndSwap(y, z)
	fmt.Println(swapped, v) // true {{7 8 9}}
}
