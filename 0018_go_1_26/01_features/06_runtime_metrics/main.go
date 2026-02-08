package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
	"sync"
	"time"
)

func main() {
	mx := new(sync.Mutex)
	go func() {
		mx.Lock()
		defer mx.Unlock()

		time.Sleep(time.Second)
	}()

	go func() {
		mx.Lock()
		defer mx.Unlock()
		time.Sleep(time.Second)
		runtime.GC()
	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Goroutine metrics:")
	printMetric("/sched/goroutines-created:goroutines", "Created")
	printMetric("/sched/goroutines:goroutines", "Live")
	printMetric("/sched/goroutines/not-in-go:goroutines", "Syscall/CGO")
	printMetric("/sched/goroutines/runnable:goroutines", "Runnable")
	printMetric("/sched/goroutines/running:goroutines", "Running")
	printMetric("/sched/goroutines/waiting:goroutines", "Waiting")

	fmt.Println("Thread metrics:")
	printMetric("/sched/gomaxprocs:threads", "Max")
	printMetric("/sched/threads/total:threads", "Live")
}

func printMetric(name string, description string) {
	sample := []metrics.Sample{
		{
			Name: name,
		},
	}

	metrics.Read(sample)
	fmt.Printf("  %s: %v\n", description, sample[0].Value.Uint64())
}
