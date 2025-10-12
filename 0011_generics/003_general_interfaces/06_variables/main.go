package main

// Interfaces that are not basic may only be used as type constraints, or as elements of other interfaces
// used as constraints. They cannot be the types of values or variables, or components of other, non-interface types.

//var x Float // illegal: Float is not a basic interface
//
//var x interface{} = Float(nil) // illegal
//
//type Floatish struct {
//	f Float // illegal
//}
//
//type Float interface {
//	~float32 | ~float64
//}
