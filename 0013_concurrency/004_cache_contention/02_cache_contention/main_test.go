package main

import (
	"math/rand/v2"
	"runtime"
	"sync/atomic"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sys/cpu"
)

const cacheLineSize = unsafe.Sizeof(cpu.CacheLinePad{})

func BenchmarkCacheContention(b *testing.B) {
	const iterations = 100
	workers := runtime.GOMAXPROCS(0)

	b.Run("with contention", func(b *testing.B) {
		type data1 struct {
			value atomic.Int64
		}

		g := new(errgroup.Group)
		g.SetLimit(workers)
		s := make([]data1, workers)

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				g.Go(func() error {
					for range iterations {
						index := rand.N[int](len(s))
						s[index].value.Add(1)
					}

					return nil
				})
			}
		})

		err := g.Wait()
		require.NoError(b, err)
	})

	b.Run("without contention", func(b *testing.B) {
		type data2 struct {
			value atomic.Int64
			_     [cacheLineSize - unsafe.Sizeof(atomic.Int64{})]byte
		}

		g := new(errgroup.Group)
		g.SetLimit(workers)
		s := make([]data2, workers)

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				g.Go(func() error {
					for range iterations {
						index := rand.N[int](len(s))
						s[index].value.Add(1)
					}

					return nil
				})
			}
		})

		err := g.Wait()
		require.NoError(b, err)
	})
}
