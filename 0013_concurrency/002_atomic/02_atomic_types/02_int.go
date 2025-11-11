package main

import "sync/atomic"

func main() {
	_ = atomic.Int32{}
	_ = atomic.Int64{}
	_ = atomic.Uint32{}
	_ = atomic.Uint64{}
	_ = atomic.Uintptr{}
}
