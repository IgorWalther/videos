package main

import "fmt"

func main() {
	var myTypePointer *MyType
	fmt.Println(myTypePointer == nil) // true
	fmt.Println(myTypePointer.Foo())
}

type MyType struct {
}

func (m *MyType) Foo() string {
	return "Hello"
}
