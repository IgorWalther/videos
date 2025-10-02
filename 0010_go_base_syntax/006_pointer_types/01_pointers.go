package main

import "fmt"

func main() {
	value := 1
	pointer := &value

	changeValue(pointer)
	fmt.Println(value)    // 2
	fmt.Println(*pointer) // 2

	changePointer(&pointer)
	fmt.Println(value)    // 2
	fmt.Println(*pointer) // 0

	var _ *************int
}

func changeValue(p *int) {
	*p = 2
}

func changePointer(p **int) {
	*p = new(int)
}
