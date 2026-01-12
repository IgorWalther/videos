package main

import (
	"fmt"
	"reflect"
)

func main() {
	err := bar()
	fmt.Println(err == nil)                   // false
	fmt.Println(reflect.TypeOf(err))          // *main.MyError
	fmt.Println(reflect.ValueOf(err).IsNil()) // true

}

type MyError struct{}

func (m *MyError) Error() string {
	return "test"
}

func bar() error {
	var m *MyError

	return m
}
