package concurrency

import (
	"context"
)

func Generator[T any](ctx context.Context, data []T, size int) <-chan T {
	result := make(chan T, size)

	go func() {
		defer close(result)

		for i := 0; i < len(data); i++ {
			select {
			case result <- data[i]:
			case <-ctx.Done():
				return
			}

		}
	}()

	return result
}
