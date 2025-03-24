package main

import (
	concurrency "concurrency_patterns"
	"context"
	"fmt"
	"net/http"
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

	urlInput := concurrency.Generator(ctx, urls, channelSize)
	resCh := concurrency.Start(ctx, workerCount, urlInput, func(currentUrl string) result {
		response, err := http.Get(currentUrl)

		if err != nil {
			return result{
				url: currentUrl,
				err: fmt.Errorf("failed %s, error - %v", currentUrl, err),
			}
		}

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
