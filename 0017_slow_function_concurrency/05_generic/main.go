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

	result, err := run[int](ctx)
	fmt.Println(result, err)
}

var (
	ErrTimeoutExceeded = errors.New("timeout")
	ErrInvariantError  = errors.New("invariant error")
)

func run[T any](ctx context.Context) (T, error) {
	resCh := make(chan T)
	errCh := make(chan error, 1)

	go func() {
		defer close(resCh)
		defer close(errCh)

		select {
		case v, ok := <-slowFunction[T]():
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

func slowFunction[T any]() <-chan T {
	ch := make(chan T, 1)

	go func() {
		defer close(ch)
		time.Sleep(time.Millisecond * rand.N[time.Duration](5000))

		ch <- *new(T)
	}()

	return ch
}
