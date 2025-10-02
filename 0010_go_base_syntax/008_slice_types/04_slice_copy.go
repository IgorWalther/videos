package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := make([]int64, 1000, 2000)
	b := make([]int64, 1000)

	copyA := intCopy(a)
	copyB := templateCopy(b)

	fmt.Println(reflect.DeepEqual(a, copyA))
	fmt.Println(reflect.DeepEqual(b, copyB))

	c := make([]int64, 0, 1000)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	c = append(c, 1)
	fmt.Println(c)
	
	s := make([]int64, 5, 10)
	t := s[2:5]

	fmt.Println(cap(t))
	fmt.Println(len(t))
}

func intCopy(source []int64) []int64 {
	dst := make([]int64, len(source))
	copy(dst, source)

	return dst
}

func templateCopy[S ~[]E, E any](source S) []E {
	dst := make([]E, len(source))
	copy(dst, source)

	return dst
}
