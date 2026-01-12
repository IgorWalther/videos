package main

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrPermission = errors.New("permission denied")
)

func readFile(name string) (string, error) {
	if name == "" {
		return "", ErrNotFound
	}

	if name == "secret.txt" {
		return "", fmt.Errorf("access to %q: %w", name, ErrPermission)
	}

	return "", fmt.Errorf("reading %q: %w", name, io.EOF)
}

// todo: clean arch layers

func main() {
	files := []string{"", "secret.txt", "data.txt"}
	for _, f := range files {
		content, err := readFile(f)

		if err != nil {
			switch {
			case errors.Is(err, ErrNotFound):
				fmt.Println(f, "file not found...")
			case errors.Is(err, ErrPermission):
				fmt.Println(f, "access denied...")
			case errors.Is(err, io.EOF):
				fmt.Println(f, "OK")
			default:
				fmt.Println(f, "undefined error:", err)
			}

			continue
		}

		fmt.Println("OK:", content)
	}
}
