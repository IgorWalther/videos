package main

import "fmt"

// A field declared with a type but no explicit field name is called an embedded field.
// An embedded field must be specified as a type name T or as a pointer to a non-interface type name *T.

type T1 struct {
	x, y int
	u    float32
	_    float32
	A    *[]int
	F    func()
}

type T2 struct {
	x, y int
}

// field names must be unique in a struct type:

type E2 struct {
	T1
	x int
	//*T1
}

func main() {
	v := E2{}
	v.x = 5

	fmt.Println(v.x)
	fmt.Println(v.T1.x)
}
