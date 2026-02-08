package main

import "testing"

var keep *byte

func benchmarkAlloc(b *testing.B, size int) {
	b.ReportAllocs()
	for b.Loop() {
		obj := make([]byte, size)
		keep = &obj[0]
	}
}

func BenchmarkAlloc1(b *testing.B) {
	benchmarkAlloc(b, 1)
}

func BenchmarkAlloc8(b *testing.B) {
	benchmarkAlloc(b, 8)
}

func BenchmarkAlloc64(b *testing.B) {
	benchmarkAlloc(b, 64)
}

func BenchmarkAlloc128(b *testing.B) {
	benchmarkAlloc(b, 128)
}

func BenchmarkAlloc512(b *testing.B) {
	benchmarkAlloc(b, 512)
}
