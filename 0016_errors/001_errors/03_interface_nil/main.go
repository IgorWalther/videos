package main

import "fmt"

func main() {
	err := bar()
	fmt.Println(err == nil) // ?
}

type MyError struct{}

func (m *MyError) Error() string {
	return "test"
}

func bar() error {
	var m *MyError

	return m
}
