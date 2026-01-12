package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := div(42, 0)

	if err != nil {
		// action
	}

	fmt.Println(res)
}

// TODO: 3 return arguments (rightmost is error)

// Your caller will check the second value to see if an error occurred

var ErrDivisionByZero = errors.New("division by zero")

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	return a / b, nil
}
