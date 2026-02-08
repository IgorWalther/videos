package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	typ := reflect.TypeFor[http.Transport]()
	for i := range typ.NumField() {
		field := typ.Field(i)
		fmt.Println(field.Name, field.Type)
	}
}
