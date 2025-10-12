package main

import (
	"cmp"
	"fmt"
)

type Response struct {
	body string
}

func main() {
	r := &Response{
		body: "123",
	}

	fmt.Println(cmp.Or(nil, nil, r))
}
