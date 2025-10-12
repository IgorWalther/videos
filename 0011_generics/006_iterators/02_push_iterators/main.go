package main

import (
	"fmt"
	"iter"
)

func Forward[T any](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for v := range Forward(nums) {
		fmt.Println("value:", v)
	}

	for v := range Forward(nums) {
		if v > 3 {
			break
		}

		fmt.Println("<=3:", v)
	}
}
