package main

type MyInt int
type MyInt2 MyInt

type (
	// I1
	// An interface representing only the type int.
	I1 interface {
		int
	}

	// I2
	// An interface representing all types with underlying type int.
	// For example, type MyInt int
	I2 interface {
		~int
	}

	// I3
	// An interface representing all types with underlying type int that implement the String method.
	I3 interface {
		~int
		String() string
	}

	// I4
	// An interface representing an empty type set: there is no type that is both an int and a string.
	I4 interface {
		int
		string
	}
)
