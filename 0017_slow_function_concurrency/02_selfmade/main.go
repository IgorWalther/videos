package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})

	wg := new(sync.WaitGroup)
	wg.Go(func() {
		time.Sleep(time.Second)
		close(done)
	})

	result, err := run(done)
	fmt.Println(result, err)

	wg.Wait()
}

var (
	ErrTimeoutExceeded = errors.New("timeout")
	ErrInvariantError  = errors.New("invariant error")
)

func run(done chan struct{}) (int, error) {
	resCh := make(chan int)
	errCh := make(chan error, 1)

	go func() {
		defer close(resCh)
		defer close(errCh)

		select {
		case v, ok := <-slowFunction():
			if !ok {
				errCh <- ErrInvariantError
				return
			}

			resCh <- v
		case <-done:
			errCh <- ErrTimeoutExceeded
			return
		}
	}()

	return <-resCh, <-errCh // left-to-right order, so resCh is non-buffered
}

func slowFunction() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		time.Sleep(time.Millisecond * rand.N[time.Duration](5000))

		ch <- 123
	}()

	return ch
}
