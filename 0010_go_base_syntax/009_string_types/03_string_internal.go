package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	stringHeader := &reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(myMalloc(100))), // not safe
		Len:  100,
	}

	s := *(*string)(unsafe.Pointer(stringHeader))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func myMalloc(size int) *byte {
	return unsafe.SliceData(make([]byte, size))
}
