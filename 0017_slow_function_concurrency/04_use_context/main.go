package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := run(ctx)
	fmt.Println(result, err)
}

var (
	ErrTimeoutExceeded = errors.New("timeout")
	ErrInvariantError  = errors.New("invariant error")
)

func run(ctx context.Context) (int, error) {
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
		case <-ctx.Done():
			errCh <- ErrTimeoutExceeded
			return
		}
	}()

	return <-resCh, <-errCh // left-to-right order, so resCh is non-buffered
}

func slowFunction() <-chan int {
	ch := make(chan int, 1)

	go func() {
		defer close(ch)
		time.Sleep(time.Millisecond * rand.N[time.Duration](5000))

		ch <- 123
	}()

	return ch
}
