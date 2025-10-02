package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println("Iteration:", i)
	}
}
