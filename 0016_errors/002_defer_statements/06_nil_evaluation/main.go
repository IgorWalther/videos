package main

import "fmt"

type Dog interface {
	Bark() string
}

type Shepherd struct {
}

func (i *Shepherd) Bark() string {
	return "Woof"
}

func main() {
	var dog Dog

	defer dog.Bark()
	dog = &Shepherd{}

	fmt.Println(dog.Bark())
}

// If a deferred function value evaluates to nil,
// execution panics when the function is invoked, not when the "defer" statement is executed.
