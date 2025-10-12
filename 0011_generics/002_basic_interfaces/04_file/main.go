package main

import (
	"fmt"
	"io/fs"
)

func main() {
	var f fs.File // nil
	fmt.Println(f)
}
