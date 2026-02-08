package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("Hello, @igoroutine")

	f, err := buf.Peek(5)
	fmt.Println(string(f), err)
	fmt.Println(string(buf.Next(5)))

	buf.Next(2)

	s, err := buf.Peek(12)
	fmt.Println(string(s), err)
}
