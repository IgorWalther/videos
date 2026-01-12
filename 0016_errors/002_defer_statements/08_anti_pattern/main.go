package main

import "fmt"

func main() {
	for i := range 100 {
		defer fmt.Println(i)
	}
}

// Possible resource leak, 'defer' is called in the 'for' loop
// defer + recover | restreamer example
