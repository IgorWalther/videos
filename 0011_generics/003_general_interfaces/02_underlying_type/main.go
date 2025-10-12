package main

type MyInt int
type MyInt2 MyInt

// In a term of the form ~T, the underlying type of T must be itself, and T cannot be an interface.

type I interface {
	~[]byte // the underlying type of []byte is itself

	//~MyInt  // illegal: the underlying type of MyInt is not MyInt
	//~error  // illegal: error is an interface
}
