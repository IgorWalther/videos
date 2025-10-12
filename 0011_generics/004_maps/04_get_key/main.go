package main

import "fmt"

func main() {
	m := make(map[string]int, 10)
	m["First"] = 1

	n := m["First"] // 1
	fmt.Println(n)

	z := m["Another"] // *new(int)
	fmt.Println(z)

	// recommended
	if v, ok := m["Another2"]; ok {
		// action
	} else {
		println(v)
	}
}
