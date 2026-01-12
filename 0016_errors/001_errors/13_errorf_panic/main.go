package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Errorf("%w %w", errors.New("1"), errors.New("2"))
}
