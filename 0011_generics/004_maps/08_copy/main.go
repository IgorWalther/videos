package main

import "maps"

func main() {
	m1 := map[string]int{
		"First": 1,
	}

	m2 := map[string]int{
		"Second": 1,
	}

	maps.Copy(m1, m2)
}
