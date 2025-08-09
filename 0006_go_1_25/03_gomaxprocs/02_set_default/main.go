package main

import "runtime"

// Sharded data structures by GOMAXPROCS ???

func main() {
	runtime.GOMAXPROCS(10)         // disable container-aware behaviour
	runtime.SetDefaultGOMAXPROCS() // enable container-aware behaviour
}
