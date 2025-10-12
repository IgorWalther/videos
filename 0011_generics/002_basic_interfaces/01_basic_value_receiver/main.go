package main

// The method set of a defined type T consists of all methods declared with receiver type T.
// The method set of a pointer to a defined type T is the set of all methods declared with receiver *T or T.

// duck typing
// interface is a set of methods

type Dog interface {
	Bark() string
}

type Shepherd struct{}

func (s Shepherd) Bark() string {
	return "Gav!"
}

type Labrador struct{}

func (l Labrador) Bark() string {
	return "Woof!"
}

func main() {
	useDog(Shepherd{}) // &Shepherd{}
	useDog(Labrador{}) // &Labrador{}
}

func useDog(dog Dog) {
	dog.Bark()
}
