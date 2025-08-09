package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	ctx := context.Background()
	fr := trace.NewFlightRecorder(trace.FlightRecorderConfig{
		MaxBytes: 5 << 20,
		MinAge:   time.Minute,
	})
	defer fr.Stop()

	if err := fr.Start(); err != nil {
		log.Fatalf("failed to start FlightRecorder: %v\n", err)
	}

	fmt.Println("Flight recorder started. Doing work...")

	ctx, task := trace.NewTask(ctx, "demo")
	defer task.End()

	go func() {
		for {
			trace.WithRegion(ctx, "time sleep", func() {
				time.Sleep(500 * time.Millisecond)
			})

			trace.WithRegion(ctx, "buffer allocation", func() {
				b := make([]byte, 100)
				b[0] = 27
			})
		}
	}()

	time.Sleep(5 * time.Second)

	f, err := os.Create("./trace.out")
	if err != nil {
		log.Fatalf("failed to create trace file: %v\n", err)
	}
	defer f.Close()

	if _, err = fr.WriteTo(f); err != nil {
		log.Fatalf("failed to write trace: %v\n", err)
	}

	fmt.Println("Trace snapshot saved to trace.out")
}
