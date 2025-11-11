package main

import (
	"testing"
)

func BenchmarkCacheContention(b *testing.B) {
	b.Run("non-atomic increment", func(b *testing.B) {
		b.ReportAllocs()

		var v int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				v++
			}
		})
	})

	b.Run("atomic increment", func(b *testing.B) {
		b.ReportAllocs()

		var v int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				SyncAdd(&v, 1)
			}
		})
	})
}
