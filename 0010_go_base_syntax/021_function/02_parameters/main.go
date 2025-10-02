package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("2 + 3 =", add(2, 3))
}
