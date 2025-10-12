package main

func main() {
	m := make(map[string]int, 10)
	m["First"] = 1
	m["Second"] = 2
	println(len(m)) // 2

	clear(m)

	println(len(m)) // 0

}
