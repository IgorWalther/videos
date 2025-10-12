package main

// The type T in a term of the form T or ~T cannot be a type parameter,
// and the type sets of all non-interface terms must be pairwise disjoint
// (the pairwise intersection of the type sets must be empty). Given a type parameter P:

type MyInt int

type I interface {
	//P               // illegal: P is a type parameter
	//int | ~P        // illegal: P is a type parameter

	// ./main.go:13:9: overlapping terms main.MyInt and ~int
	~int | MyInt // illegal: the type sets for ~int and MyInt are not disjoint (~int includes MyInt)

	float32 | Float // overlapping type sets but Float is an interface
}

type Float interface {
	~float32 | ~float64
}

func main() {}
