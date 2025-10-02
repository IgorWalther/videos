package main

import "fmt"

func main() {
	s := "Hello"
	for i, r := range s {
		fmt.Printf("Index %d, Rune %c\n", i, r)
	}
}
