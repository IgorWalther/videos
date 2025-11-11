package _3_bullshit

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkIncrement(b *testing.B) {
	b.Run("sequential increment", func(b *testing.B) {
		b.ReportAllocs()

		v := 0
		for b.Loop() {
			for range 1000 {
				v++
			}
		}
	})

	b.Run("parallel increment", func(b *testing.B) {
		b.ReportAllocs()

		v := atomic.Int64{}
		wg := new(sync.WaitGroup)

		for b.Loop() {
			for range 1000 {
				wg.Go(func() {
					v.Add(1)
				})
			}

			wg.Wait()
		}
	})
}
