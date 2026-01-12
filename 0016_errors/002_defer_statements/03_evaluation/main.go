package main

import "fmt"

type User struct {
	Age int
}

func main() {
	d := User{}

	defer fmt.Println(d.Age)
	d.Age++

	return
}

// The function value and parameters to the call are evaluated as usual and
// saved anew but the actual function is not invoked
