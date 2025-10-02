package main

// A struct is a sequence of named elements, called fields, each of which has a name and a type.
// Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).
// Within a struct, non-blank field names must be unique.

// Z
// An empty struct.
var Z struct{}

// T
// Struct with 6 fields.
type T1 struct {
	x, y int
	u    float32
	_    float32
	A    *[]int
	F    func()
}

type T2 struct {
	T1
}
