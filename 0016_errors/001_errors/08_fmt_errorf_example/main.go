package main

import (
	"errors"
	"fmt"
	"os"
)

func simulateFileError() error {
	return os.ErrNotExist
}

func readConfig() error {
	err := simulateFileError()

	if err != nil {
		return fmt.Errorf("read config failed: %w", err)
	}

	return nil
}

// > 1 %w -> Unwrap with []error

func main() {
	err := readConfig()
	if err != nil {
		fmt.Println("Full error:", err)

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist — caught via errors.Is")
		}

		var pathErr *os.PathError
		if errors.As(err, &pathErr) { // see internal
			fmt.Println("Underlying PathError:", pathErr)
		}

		fmt.Println("Unwrapped once:", errors.Unwrap(err))
	}
}
