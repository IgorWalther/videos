package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	for i, e := range slices.Backward(s) {
		fmt.Println(i, e)
	}
}
