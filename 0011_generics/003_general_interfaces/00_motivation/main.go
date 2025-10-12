package main

import (
	"cmp"
	"fmt"
)

// 1.18 (+- 2022)

func minInt(l, r int) int {
	if l < r {
		return l
	}

	return r
}

func minInt64(l, r int64) int64 {
	if l < r {
		return l
	}

	return r
}

func minFloat32(l, r float32) float32 {
	if l < r {
		return l
	}

	return r
}

func minFloat64(l, r float64) float64 {
	if l < r {
		return l
	}

	return r
}

func main() {
	fmt.Println(min(1, 2))
	fmt.Println(min(float64(1), float64(2)))
	fmt.Println(myMin(float64(1), float64(2)))
	fmt.Println(myMin(1, 2))
}

func myMin[T cmp.Ordered](l, r T) T {
	if l < r {
		return l
	}

	return r
}
