package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

type config struct {
}

var curCfg *config

func updater() {
	for {
		time.Sleep(time.Second)

		oldPtr := unsafe.Pointer(curCfg)
		newCfg := config{}

		atomic.SwapPointer(&oldPtr, unsafe.Pointer(&newCfg))
	}
}

func main() {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old) // 789 123
}
