package main

import "fmt"

// duck typing

type Student interface {
	Code()
}

type Dog interface {
	Bark() string
}

type TrainedDog interface {
	Bark() string
	Sit() string
	Fetch(item string) string
	Stay() string
}

type Shepherd struct{}

func (s *Shepherd) Bark() string {
	return "Woof!"
}

func (s *Shepherd) Sit() string {
	return "Shepherd sits down."
}

func (s *Shepherd) Fetch(item string) string {
	return fmt.Sprintf("Shepherd fetches the %s.", item)
}

func (s *Shepherd) Stay() string {
	return "Shepherd stays still."
}

type Labrador struct{}

func (l *Labrador) Bark() string {
	return "Woof!"
}

func main() {
	shepherd := &Shepherd{}
	dog := trainedDogToDog(shepherd)
	fmt.Println(dog.Bark())

	//trainedDogToStudent(shepherd)
}

// panic: interface conversion: *main.Shepherd is not main.Student: missing method Code
// A.(B) - ok, if B methods in A
func trainedDogToDog(trainedDog TrainedDog) Dog {
	//v := trainedDog.(Dog)

	v, ok := trainedDog.(Dog)

	if !ok {
		panic("cast error")
	}

	return v
}

func trainedDogToStudent(trainedDog TrainedDog) Student {
	return trainedDog.(Student) // panic
}
