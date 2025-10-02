package main

import "fmt"

// An array is a numbered sequence of elements of a single type, called the element type.
// The number of elements is called the length of the array and is never negative.

// The length is part of the array's type;
// it must evaluate to a non-negative constant representable by a value of type int
// The length of array a can be discovered using the built-in function len.

func main() {
	var _ [5]int

	_ = [5]int{}
	_ = [5]int{1, 2, 3, 4, 5}
	_ = [5]int{2: 5} // [0,0,5,0,0]

	array := [1024]int{}
	fmt.Println(len(array))
	fmt.Println(cap(array))

	functionWithArray(&array) // copy here

	// The elements can be addressed by integer indices 0 through len(a)-1.
	fmt.Println(array[0]) // 0
}

func functionWithArray(array *[1024]int) {
	array[0] = 42
}

// Array types are always one-dimensional but may be composed to form multi-dimensional types.
// [32]byte
// [2*N] struct { x, y int32 }
// [1000]*float64
// [3][5]int
// [2][2][2]float64  // same as [2]([2]([2]float64))
