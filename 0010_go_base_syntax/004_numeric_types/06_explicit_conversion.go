package main

import "fmt"

func main() {
	//var n int64 = 5
	// foo(n) - compilation error

	foo(5) // auto literal cast
}

// To avoid portability issues all numeric types are defined types and thus distinct except byte,
// which is an alias for uint8, and rune, which is an alias for int32. Explicit conversions are required when
// different numeric types are mixed in an expression or assignment. For instance, int32 and int are not the same
// type even though they may have the same size on a particular architecture.
func foo(v int) {
	fmt.Println(v)
}
