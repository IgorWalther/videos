package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	for v := range maps.Values(m) {
		fmt.Println(v)
	}

	// maps.Insert()
	// maps.Collect()
}
