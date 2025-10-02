package main

import (
	"fmt"
)

// A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
// The number of bytes is called the length of the string and is never negative. Strings are immutable:
// once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type.

// The length of a string s can be discovered using the built-in function len.
// The length is a compile-time constant if the string is a constant.
// A string's bytes can be accessed by integer indices 0 through len(s)-1. It is illegal to take the address of such an element;
// if s[i] is the i'th byte of a string, &s[i] is invalid.

func main() {
	fmt.Println(`abc`)                 // same as "abc"
	fmt.Println(`Привет, 123 123 123`) // same as "\\n\n\\n"
	fmt.Println("\n")
	fmt.Println("\"") // same as `"`
	fmt.Println("Hello, world!\n")
	fmt.Println("日本語")
	fmt.Println("\u65e5本\U00008a9e")
	fmt.Println("\xff\u00FF")

	s := "日本語"
	fmt.Println(len(s))
	fmt.Println(s[2])
}
