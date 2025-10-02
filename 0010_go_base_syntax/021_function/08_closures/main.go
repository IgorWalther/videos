package main

import "fmt"

func counter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	c := counter()
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3
}
