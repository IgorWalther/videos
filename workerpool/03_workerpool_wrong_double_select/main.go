package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func Start(
	ctx context.Context,
	workersCount int,
	input <-chan int,
	transform func(e int) int,
) <-chan int {
	result := make(chan int)
	wg := new(sync.WaitGroup)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-input:
					if !ok {
						return
					}

					result <- transform(v)
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int, 10)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	close(ch)

	Start(ctx, 10, ch, func(e int) int {
		return e + 1
	})
	cancel()

	for range 1000000 {
		fmt.Println(runtime.NumGoroutine())
	}
}
