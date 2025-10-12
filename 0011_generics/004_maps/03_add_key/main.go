package main

func main() {
	m := make(map[string]int, 10)
	m["First"] = 1
	m["Second"] = 2
	m["Third"] = 3

	_ = map[string]int{
		"First":  1,
		"Second": 2,
		"Third":  3,
	}
}
