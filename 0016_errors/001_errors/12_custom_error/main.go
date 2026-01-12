package main

import (
	"fmt"
)

type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s %q not found", e.Resource, e.ID)
}

var ErrNotFound = &NotFoundError{}

func (e *NotFoundError) Is(target error) bool {
	_, ok := target.(*NotFoundError)
	return ok
}

// if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
//			return true
//		}

func main() {
	// errors.Is(err, ErrNotFound)
}
