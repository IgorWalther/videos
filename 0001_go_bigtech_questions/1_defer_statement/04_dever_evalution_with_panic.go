package main

import "fmt"

// While executing a function F, an explicit call to panic or a run-time panic
// terminates the execution of F. Any functions deferred by F are then executed
// as usual. Next, any deferred functions run by F's caller are run, and so on up
// to any deferred by the top-level function in the executing goroutine.
// At that point, the program is terminated and the error condition is reported,
// including the value of the argument to panic
func main() {
	defer func() {
		fmt.Println("defer2")
	}()

	F()
}

func F() {
	defer func() {
		fmt.Println("defer1")
	}()

	panic("panic")
}
