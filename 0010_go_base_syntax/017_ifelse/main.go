package main

import (
	"fmt"
	"os"
)

func main() {
	var x int
	fmt.Fscan(os.Stdin, &x)

	if x > 0 {
		fmt.Println("positive")
	} else if x < 0 {
		// ...
	} else {
		// ...
	}

	if x := foo(); x > 10 {
		// ...
	}
}

func foo() int {
	return 5
}
