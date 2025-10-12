package main

import "maps"

func main() {
	m := make(map[string]int, 10)
	m["First"] = 1

	// Idempotent operation
	delete(m, "First")
	delete(m, "First")
	delete(m, "First")

	maps.DeleteFunc(m, func(k string, v int) bool {
		return k == "First" && v != 0
	})
}
