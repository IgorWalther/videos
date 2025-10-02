package main

import (
	"fmt"
	"strings"
)

func main() {
	s := &strings.Builder{}

	for i := 0; i < 100; i++ {
		s.WriteByte((byte)(i))
	}

	fmt.Println(s.String())
}
