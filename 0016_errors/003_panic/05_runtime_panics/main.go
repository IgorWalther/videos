package main

import "runtime"

// The Error interface identifies a run time error.
type Error interface {
	error

	// RuntimeError is a no-op function but
	// serves to distinguish types that are run time
	// errors from ordinary errors: a type is a
	// run time error if it has a RuntimeError method.
	RuntimeError()

	// and perhaps other methods
}

func main() {
	_ = runtime.PanicNilError{}
	_ = runtime.TypeAssertionError{} // i.(T)
}

// Execution errors such as attempting to index an array out of bounds trigger a run-time panic equivalent to a call
// of the built-in function panic with a value of the implementation-defined interface type runtime.Error.
// That type satisfies the predeclared interface type error. The exact error values that represent distinct
// run-time error conditions are unspecified.
