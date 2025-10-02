package main

import (
	"fmt"
	"os"
)

func main() {
	var x int
	fmt.Fscan(os.Stdin, &x)

	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
		fallthrough
	case x > 0 && x <= 10:
		fmt.Println("x is between 1 and 10")
	default:
		fmt.Println("x is greater than 10")
	}
}
