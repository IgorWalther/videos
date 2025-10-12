package main

// any is an alias for interface{} and is equivalent to interface{} in all ways.
//type any = interface{}

// Every type that is a member of the type set of an interface implements that interface.
// Any given type may implement several distinct interfaces.
// For instance, all types implement the empty interface which stands for the set of all (non-interface) types:

// func foo (v interface{})
func foo(v any) {
	// v.()
}

func main() {
	foo(1)
	foo(2)

	type dog struct {
		name string
	}

	foo(dog{})
}
