package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex
var v string

func f() {
	v = "the nature of concurrency"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()

	fmt.Println(v)
}
