package main

import (
	"fmt"
)

// A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence
// of elements from that array. A slice type denotes the set of all slices of arrays of its element type.
// The number of elements is called the length of the slice and is never negative.
// The value of an uninitialized slice is nil.

func main() {
	var _ []int // zero value - nil
	_ = []int{} // []

	s := make([]int, 5, 5)
	fmt.Println(s)
}
