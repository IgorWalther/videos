package main

import (
	"sync"
)

func main() {
	// See Mullender and Cox, ``Semaphores in Plan 9,''
	_ = sync.WaitGroup{}
	_ = sync.Mutex{}
}
