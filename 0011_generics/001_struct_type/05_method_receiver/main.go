package main

// The method set of a defined type T
// consists of all methods declared with receiver type T.

// The method set of a pointer to a defined type T
// is the set of all methods declared with receiver *T or T.

// duck typing
// interface is a set of methods

type Dog struct {
	name string
}

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

func main() {
	d := Dog{}
	d.Bark()
	d.Name()

	dPointer := &Dog{}
	dPointer.Bark()
	dPointer.Name()
}

// (&d).Bark() !!!
