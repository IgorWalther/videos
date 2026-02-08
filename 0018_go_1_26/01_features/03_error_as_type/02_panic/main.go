package main

import (
	"errors"
	"fmt"
)

type AppError struct {
	Msg string
}

func makeErr() error {
	return fmt.Errorf("level1: %w", fmt.Errorf("level2: %w", &AppError{
		Msg: "database is down",
	}))
}

func main() {
	err := makeErr()
	// var target *AppError

	var target *AppError

	// panic: errors: *target must be interface or implement error
	// if errors.As(err, &target) {
	//	fmt.Println("As:", target.Msg)
	// }

	// panic: errors: target cannot be nil
	// if errors.As(err, &target) {
	//	fmt.Println("As:", target.Msg)
	// }

	if errors.As(err, nil) {
		fmt.Println("As:", target.Msg)
	}

	// if _, ok := errors.AsType[*AppError](err); ok {}
}
