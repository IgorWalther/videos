package main

import (
	"fmt"
)

// The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.
func main() {
	fmt.Printf("%08b\n", uint8(getTwo()))
	fmt.Printf("%08b\n", 2)
}

// go:noinline
func getTwo() int8 {
	return -2
}
