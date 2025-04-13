package main

import "fmt"

func main() {
	m := map[string]any{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": "undefined",
	}

	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}
