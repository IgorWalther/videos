package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	b := atomic.Bool{}
	b.Store(true)

	fmt.Println(b.Load())
}
