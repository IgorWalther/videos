package main

import "fmt"

type Base struct {
	A int
}

func (b *Base) Speak() {
	fmt.Println("Base speaks!")
}

type Derived struct {
	Base
	B int
}

func main() {
	d := &Derived{}
	fmt.Println(d.A)
	fmt.Println(d.B)
	fmt.Println(d.Base.A)
	d.Speak() // Derived speaks!
}

func (d *Derived) Speak() {
	fmt.Println("Derived speaks!")
}

// Given a struct type S and a named type T, promoted methods are included in the method set of the struct as follows:

// If S contains an embedded field T, the method sets of S and *S both include promoted methods with receiver T.
// The method set of *S also includes promoted methods with receiver *T.

// If S contains an embedded field *T, the method sets of S and *S both include promoted methods with receiver T or *T.
