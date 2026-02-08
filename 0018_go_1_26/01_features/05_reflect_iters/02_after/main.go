package main

import (
	"fmt"
	"net/http"
	"reflect"
)

// The new methods Type.Fields, Type.Methods, Type.Ins and Type.Outs return iterators for a type’s fields (for a struct type),
// methods, inputs and outputs parameters (for a function type), respectively.

func main() {
	typ := reflect.TypeFor[http.Client]()
	for f := range typ.Fields() {
		fmt.Println(f.Name, f.Type)
	}
}
