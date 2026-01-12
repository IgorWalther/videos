package main

import (
	"fmt"

	"github.com/pkg/errors" // extensions
)

func main() {
	err := foo()

	fmt.Println(errors.Is(err, ErrExample))

	// errors.Unwrap()
	// errors.WithStack()
	// errors.Cause()
	// errors.WithMessage()
}

var ErrExample = errors.New("test")

func foo() error {
	return errors.Wrap(ErrExample, "message")
}
