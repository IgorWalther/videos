package main

import (
	"sync"
	"testing"
)

func BenchmarkMutexAdd(b *testing.B) {
	b.ReportAllocs()

	var number int32
	mutex := new(sync.Mutex)

	b.ResetTimer()

	for b.Loop() {
		mutex.Lock()
		number++
		mutex.Unlock()
	}
}

func BenchmarkRWMutexAdd(b *testing.B) {
	b.ReportAllocs()

	var number int32
	mutex := new(sync.RWMutex)

	b.ResetTimer()

	for b.Loop() {
		mutex.Lock()
		number++
		mutex.Unlock()
	}
}
