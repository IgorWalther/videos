package mutexbench

import (
	"sync"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	benchcases := []struct {
		name string
		lock sync.Locker
	}{
		{
			name: "ticket lock",
			lock: NewTicketLock(),
		},
		{
			name: "standard library mutex",
			lock: new(sync.Mutex),
		},
	}

	for _, tt := range benchcases {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetParallelism(3)

			repo := NewRepository(tt.lock)

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					repo.Operation()
				}
			})
		})
	}
}
