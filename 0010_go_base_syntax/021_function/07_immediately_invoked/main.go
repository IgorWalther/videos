package main

import "fmt"

func main() {
	result := func(x int) int {
		return x * x
	}(5)

	fmt.Println("Square:", result)
}
