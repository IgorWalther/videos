package main

import (
	"reflect"
	_ "unsafe"
)

// but avoids unnecessary memory allocations
func before[T any](rv reflect.Value) (T, bool) {
	v, ok := reflect.TypeAssert[T](rv)
	return v, ok
}

func after[T any](rv reflect.Value) (T, bool) {
	v, ok := rv.Interface().(T)
	return v, ok
}
