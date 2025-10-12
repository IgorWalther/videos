package main

type Dog struct {
	name string
}

// naming!

// Name
// func Name(d Dog) string
func (d Dog) Name() string {
	return d.name
}

// Bark
// func Bark(d *Dog) string
func (d *Dog) Bark() string {
	return "Woof!"
}
