package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sync"
	"sync/atomic"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filepath.Join(wd, "0004_profiling/01_collect_profile/cpu.out"))

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalln(err)
	}

	defer pprof.StopCPUProfile()

	wg := new(sync.WaitGroup)
	for i := 0; i < runtime.NumCPU()/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range int(10e6) {
				root()
			}
		}()
	}

	wg.Wait()
}

var v atomic.Int32

//go:noinline
func root() {
	f2()
	f3()
	f4()
}

func f2() {
	v.Add(1)
}

func f3() {
	v.Add(1)
}

func f4() {
	v.Add(1)
}
