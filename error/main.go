package main

import (
	"fmt"
	"reflect"
)

func main() {
	err := process()
	fmt.Println(err == nil)                   // false
	fmt.Println(reflect.TypeOf(err))          // *main.MyError
	fmt.Println(reflect.ValueOf(err).IsNil()) // true

	if err != nil {
		panic(err)
	}
}

type MyError struct {
}

func (m *MyError) Error() string {
	return "my_error_message"
}

func process() error {
	var m *MyError

	// if ... {
	//	  m = &MyError{}
	//}

	fmt.Println(m == nil) // true
	return m
}

//type interface struct {
//	data unsafe.Pointer
//	table _table
//}
