package main

import (
	"math/rand/v2"
	"time"
)

func main() {
	run() // 1s

	// ...
}

func run() {
	slowFunction()
}

func slowFunction() int {
	time.Sleep(time.Millisecond * rand.N[time.Duration](5000))

	return 123
}
