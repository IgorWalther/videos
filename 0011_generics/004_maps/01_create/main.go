package main

import "fmt"

func main() {
	var m map[int]int
	fmt.Println(m) // nil

	_ = make(map[int]int, 10) // recommended

	_ = map[int]int{
		1: 42,
	}

	//map[*T]struct{ x, y float64 }
	//map[string]interface{}
}
