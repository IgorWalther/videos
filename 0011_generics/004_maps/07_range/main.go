package main

import "fmt"

func main() {
	m := map[string]int{
		"First":  1,
		"Second": 2,
		"Third":  3,
	}

	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v) // ?
	}

	for k := range m {
		fmt.Printf("key: %s\n", k) // ?
	}
}
