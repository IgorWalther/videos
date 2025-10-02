package main

import (
	"fmt"
)

func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func main() {
	a, p := rectangle(5, 3)
	fmt.Println("Area:", a, "Perimeter:", p)
}
