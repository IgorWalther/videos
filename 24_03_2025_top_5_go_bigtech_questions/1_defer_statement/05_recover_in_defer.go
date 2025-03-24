package main

import "fmt"

func main() {
	protect(func() {
		panic("test\n")
	})

	fmt.Println("all done")
}

func protect(g func()) {
	defer func() {
		fmt.Println("done")

		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()

	fmt.Println("start")
	g()
}
