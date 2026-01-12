package main

import "fmt"

func main() {
	fmt.Println(named())
}

func named() (res int) {
	defer func() {
		res++
	}()

	return 42
}

// Deferred functions are executed after any result parameters are set by that
// return statement but before the function returns to its caller
