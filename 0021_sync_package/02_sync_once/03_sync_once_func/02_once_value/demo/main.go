package main

import "sync"

func main() {
	f := sync.OnceFunc(foo)
	f()
}

func foo() {
	panic("hello")
}
