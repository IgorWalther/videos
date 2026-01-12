package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

// Deferred functions are invoked immediately before the surrounding
// function returns, in the reverse order they were deferred
