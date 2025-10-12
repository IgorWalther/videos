package main

import "fmt"

// Implementation restriction: A union (with more than one term) cannot contain the predeclared
// identifier comparable or interfaces that specify methods, or embed comparable or interfaces that specify methods.

// legal
type I1 interface {
	String() string
	comparable
}

// illegal
type I2 interface {
	String() string
	//comparable | Float
}

type I3 interface {
	comparable
	String() string
}

type I4 interface {
	comparable
	fmt.Stringer
}

type Float interface {
	~float32 | ~float64
}

// legal
type I5 interface {
	comparable
}

// illegal
//type I6 interface {
//	Print(a comparable)
//}

func main() {}
