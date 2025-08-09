package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMockedTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		fmt.Println(time.Now()) // 2000-01-01 03:00:00 +0300 MSK
	})
}

func TestImmediatelyReturned(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		fmt.Println(time.Now())

		// clock moves forward instantaneously if all goroutines in the bubble are blocked
		select {
		case <-time.After(time.Hour):
			fmt.Println("1 HOUR ???")
		}
	})
}

func TestWaitGoroutines(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		gNum := runtime.NumGoroutine()
		someCode(ctx)

		synctest.Wait()

		require.Equal(t, workers+gNum, runtime.NumGoroutine())
		cancel()
		synctest.Wait()

		require.Equal(t, gNum, runtime.NumGoroutine())
	})
}

const workers = 100

func someCode(ctx context.Context) {
	for range workers {
		go func() {
			select {
			case <-ctx.Done():
				return
			}
		}()
	}
}
