package main

import (
	"fmt"
)

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case error:
				err = fmt.Errorf("panic caught: %w", x)
			default:
				err = fmt.Errorf("panic caught: %v", x)
			}

			result = 0
		}
	}()

	return a / b, nil
}

func main() {
	res, err := safeDivide(10, 0)
	fmt.Println("result:", res)

	if err != nil {
		fmt.Println("error:", err)
	}

	res, err = safeDivide(10, 2)
	fmt.Println("result:", res, "error:", err)
}
