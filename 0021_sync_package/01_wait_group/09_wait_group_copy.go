package main

import "sync"

// Package sync provides basic synchronization primitives such as mutual exclusion locks.
// Other than the Once and WaitGroup types, most are intended for use by low-level library routines.
// Higher-level synchronization is better done via channels and communication.
// Values containing the types defined in this package should not be copied.

// panic in sync.Cond

func done(wg sync.WaitGroup) {
	wg.Done() // -1
}

// all goroutines are asleep - deadlock!
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	done(wg)

	wg.Wait()
}
