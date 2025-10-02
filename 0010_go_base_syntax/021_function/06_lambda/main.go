package main

import "fmt"

func main() {
	multiply := func(a, b int) int { // first class objects
		return a * b
	}
	fmt.Println("3 * 4 =", multiply(3, 4))

	square := func(n int) int { return n * n }
	fmt.Println("Square of 6 =", apply(square, 6))

	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println("Double 4 =", double(4))
	fmt.Println("Triple 4 =", triple(4))
}

func apply(f func(int) int, x int) int {
	return f(x)
}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
