package main

import (
	"fmt"
)

func changeArray(arr [3]int) {
	arr[0] = 99
	fmt.Println("Inside function:", arr)
}

func main() {
	original := [3]int{1, 2, 3}
	changeArray(original)
	fmt.Println("In main:", original)
}
