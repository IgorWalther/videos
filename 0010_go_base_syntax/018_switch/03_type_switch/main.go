package main

import "fmt"

func printType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is an int:", v)
	case string:
		fmt.Println("This is a string:", v)
	case bool:
		fmt.Println("This is a bool:", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	printType(42)
	printType("hello")
	printType(true)
	printType(3.14)
}
