package main

import "fmt"

// go mod init (todo: workspace)
// go build
// go run main.go
// go test

func main() {
	fmt.Println(greater("World"))
}

func greater(name string) string {
	return "Hello, " + name + "!"
}
