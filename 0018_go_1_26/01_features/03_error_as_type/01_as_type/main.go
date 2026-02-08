package main

import (
	"errors"
	"fmt"
)

type AppError struct {
	Msg string
}

func (e *AppError) Error() string {
	return e.Msg
}

func makeErr() error {
	return fmt.Errorf("level1: %w", fmt.Errorf("level2: %w", &AppError{
		Msg: "database is down",
	}))
}

func main() {
	err := makeErr()

	// Go 1.13+
	var target *AppError
	if errors.As(err, &target) {
		fmt.Println("As:", target.Msg)
	}

	// Go 1.26+
	if target, ok := errors.AsType[*AppError](err); ok {
		fmt.Println("AsType:", target.Msg)
	}
}
