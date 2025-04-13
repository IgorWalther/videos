package concurrency

import (
	"context"
	"sync"
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
