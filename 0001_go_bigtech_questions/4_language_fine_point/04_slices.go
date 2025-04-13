package main

import "fmt"

func main() {
	first := make([]int, 10)
	fmt.Println(len(first)) // 10
	fmt.Println(cap(first)) // 10
	//
	//second := first[8:9]
	//fmt.Println(len(second))
	//fmt.Println(cap(second))
	//second = append(second, 5)
	//fmt.Println(first[9])

	third := first[8:9:9]
	fmt.Println(len(third))
	fmt.Println(cap(third))

	third = append(third, 5)
	fmt.Println(first[9])
}
