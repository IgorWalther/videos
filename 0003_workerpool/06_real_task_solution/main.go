package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

const (
	workerCount = 3
	channelSize = 3
)

func main() {
	ctx := context.Background()

	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	type result struct {
		url string
		err error
	}

	urlInput := Generator(ctx, urls, channelSize)
	resCh := Start(ctx, workerCount, urlInput, func(currentUrl string) result {
		response, err := http.Get(currentUrl)

		if err != nil {
			return result{
				url: currentUrl,
				err: fmt.Errorf("failed %s, error - %v", currentUrl, err),
			}
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return result{
				url: currentUrl,
				err: fmt.Errorf("failed %s with http code %d", currentUrl, response.StatusCode),
			}
		}

		return result{
			url: currentUrl,
		}
	})

	for r := range resCh {
		fmt.Printf("URL: %s, Error: %v\n", r.url, r.err)
	}
}

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
