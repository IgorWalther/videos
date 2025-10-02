package main

import "fmt"

func main() {
	fmt.Println(createZeroValue[int]())         // 0
	fmt.Println(createZeroValue[float64]())     // 0
	fmt.Println(createZeroValue[string]())      // ""
	fmt.Println(createZeroValue[interface{}]()) // nil
	fmt.Println(createZeroValue[struct{}]())    // {}
	fmt.Println(createZeroValue[[]int]())       // nil ([])
	fmt.Println(createZeroValue[*int]())        // nil
	fmt.Println(createZeroValue[[5]int]())      // [0 0 0 0 0]
	fmt.Println(createZeroValue[func()]())      // [0 0 0 0 0]
}

func createZeroValue[T any]() T {
	var t T
	return t

	// return *new(T)
}
