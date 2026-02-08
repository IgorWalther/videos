package perfbench

import "testing"

func BenchmarkCgoCall(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		cgoNop()
	}
}

func BenchmarkCgoCallParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cgoNop()
		}
	})
}

func BenchmarkCgoCallWithCallback(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		cgoCallWithCallback()
	}
}
