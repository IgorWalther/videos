package main

import "os"

func main() {
	f, err := os.Open("nonExistentFile")
	name := f.Name() // shift code to line 11
	if err != nil {
		return
	}
	println(name)
}
