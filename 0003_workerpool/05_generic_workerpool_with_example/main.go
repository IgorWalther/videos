package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Start[T, R any](
	ctx context.Context,
	workersCount int,
	input <-chan T,
	transform func(e T) R,
) <-chan R {
	result := make(chan R)
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

					select {
					case <-ctx.Done():
						return
					case result <- transform(v):
					}
				}
			}
		}()
	}

	go func() {
		defer close(result)
		wg.Wait()
	}()

	return result
}

func main() {
	ctx := context.Background()
	ch := make(chan int, 10)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 3
	ch <- 3
	ch <- 3
	ch <- 3
	ch <- 3
	ch <- 3
	ch <- 3
	close(ch)

	result := Start[int, string](ctx, 10, ch, func(e int) string {
		time.Sleep(time.Second)
		return "hello" + strconv.Itoa(e)
	})

	for e := range result {
		fmt.Println(e)
	}
}
