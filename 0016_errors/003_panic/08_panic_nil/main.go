package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		} else {
			fmt.Println("recover() saw nil panic value") // before 1.21
		}
	}()

	_ = runtime.PanicNilError{}

	panic(nil)
}
