package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// В основе многопоточных алгоритмов лежит CAS
// Консенсусное число +inf

func main() {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old) // 789 123

	swapped := atomic.CompareAndSwapInt64(&n, 123, 456)
	fmt.Println(swapped) // false
	fmt.Println(n)       // 789
	swapped = atomic.CompareAndSwapInt64(&n, 789, 456)
	fmt.Println(swapped) // true
	fmt.Println(n)       // 456
}

type Consensus[T any] struct {
	mx       sync.Mutex
	decision T
	done     bool
}

// Консенсусыв разных областях определяют чуть по-разному

// Agreement
// Validity
// Termination

func (c *Consensus[T]) NonConsistencyDecide(proposal T) T {
	return proposal
}

func (c *Consensus[T]) NonValidDecide(_ T) T {
	return *new(T)
}

func (c *Consensus[T]) NonWaitFreeDecide(proposal T) T {
	c.mx.Lock()

	if !c.done {
		c.decision = proposal
		c.done = true
	}

	c.mx.Unlock()

	return c.decision
}

// Distributed systems

// Termination
// Eventually, every correct process decides some value

// Integrity
// If all the correct processes proposed the same value v, then any correct process must decide v

// Agreement
// Every correct process must agree on the same value
