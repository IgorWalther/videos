package main

// A struct type T may not contain a field of type T,
// or of a type containing T as a component, directly or indirectly,
// if those containing types are only array or struct types

// invalid struct types
//type (
//	T1 struct{ T1 }            // T1 contains a field of T1
//	T2 struct{ f [10]T2 }      // T2 contains T2 as component of an array
//	T3 struct{ T4 }            // T3 contains T3 as component of an array in struct T4
//	T4 struct{ f [10]T3 }      // T4 contains T4 as component of struct T3 in an array
//)

// valid struct types
type (
	T5 struct{ f *T5 }       // T5 contains T5 as component of a pointer
	T6 struct{ f func() T6 } // T6 contains T6 as component of a function type
	T7 struct{ f [10][]T7 }  // T7 contains T7 as component of a slice in an array
)
