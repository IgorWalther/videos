package main

import "runtime/debug"

func main() {
	debug.SetGCPercent(100)
	debug.SetMemoryLimit(100)
}
