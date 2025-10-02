package main

import (
	"fmt"
	"unsafe"
)

// uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
func main() {
	fmt.Println(unsafe.Sizeof(uintptr(0))) // 8

	s := []int64{1, 2}
	p := &s[0]

	//next := unsafe.Add(unsafe.Pointer(p), 8)
	next := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8))

	fmt.Println(*next) // 2
}
