package main

import (
	"slices"
	"sync"
)

// before go fix
func contains(s []int, x int) bool {
	return slices.Contains(s, x)
}

func bar() interface{} {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	return 42
}

func ptrOf[T any](v T) *T {
	return &v
}
