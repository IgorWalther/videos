package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type CancelFunc = func()

var ErrContextCancelled = errors.New("context canceled")
var ErrTimeout = errors.New("timeout")

type MyContext struct {
	done chan struct{}
	err  error
	once *sync.Once
	mx   *sync.Mutex
}

func NewContext(timeout time.Duration) (*MyContext, CancelFunc) {
	ctx := &MyContext{
		done: make(chan struct{}),
		once: new(sync.Once),
		mx:   new(sync.Mutex),
	}

	if timeout <= 0 {
		panic("invariant error")
	}

	go func() {
		select {
		case <-time.After(timeout):
			ctx.cancelWithError(ErrTimeout)
		case <-ctx.done:
		}
	}()

	return ctx, func() {
		ctx.cancelWithError(ErrContextCancelled)
	}
}

func (c *MyContext) Done() <-chan struct{} {
	return c.done
}

func (c *MyContext) Err() error {
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.err
}

func (c *MyContext) Cancel() {
	c.cancelWithError(errors.New("context canceled"))
}

func (c *MyContext) cancelWithError(err error) {
	c.once.Do(func() {
		c.mx.Lock()
		defer c.mx.Unlock()
		c.err = err
		close(c.done)
	})
}

func main() {
	ctx, cancel := NewContext(time.Second)
	defer cancel()

	result, err := run(ctx)
	fmt.Println(result, err)
}

var (
	ErrTimeoutExceeded = errors.New("timeout")
	ErrInvariantError  = errors.New("invariant error")
)

func run(ctx *MyContext) (int, error) {
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
	ch := make(chan int)

	go func() {
		defer close(ch)
		time.Sleep(time.Millisecond * rand.N[time.Duration](5000))

		ch <- 123
	}()

	return ch
}
