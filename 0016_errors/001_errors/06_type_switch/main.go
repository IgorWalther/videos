package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("data.txt")

	if err != nil {
		switch v := err.(type) {
		case *os.PathError:
			fmt.Println(v.Path)
		default:
			// action
		}
	}

	f.Close() // todo: defer
}

// A library writer is free to implement this interface with a richer model under the covers,
// making it possible not only to see the error but also to provide some context
