package main

import "fmt"

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)
}
