package main

import "testing"

func Foo(optionalValue *string) {}

func TestFoo(t *testing.T) {
	v := "123"
	Foo(&v)

	// Foo(new("123")) // since Go 1.26
}
