package main

import (
	"fmt"
)

type Key struct {
	ID       int
	Callback func()
}

var globalMap = map[any]any{}

// go build main.go
func main() {
	callback := func() {
		fmt.Println("@igoroutine")
	}

	key1 := Key{
		ID:       1,
		Callback: callback,
	}

	key2 := Key{
		ID:       1,
		Callback: callback,
	}

	insert[Key](key1)
	insert[Key](key2)
	fmt.Println(len(globalMap))
}

func insert[T any](key T) {
	globalMap[key] = 27
}

// The comparison operators == and != must be fully defined for operands
// of the key type; thus the key type must not be a function, map, or slice.
// If the key type is an interface type, these comparison operators must be defined
// for the dynamic key values; failure will cause a run-time panic.
