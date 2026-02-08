package main

import (
	"encoding/json/v2"
	"fmt"
	"time"
)

type Person struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  *int   `protobuf:"bytes,2,opt,name=age,proto3" json:"age,omitempty"`
}

func personJSON(name string, born time.Time) ([]byte, error) {
	return json.Marshal(Person{
		Name: name,
		Age:  ptr(yearsSince(born)),
	})
}

func yearsSince(t time.Time) int {
	return int(time.Since(t).Hours() / (365 * 24))
}

func ptr[T any](v T) *T {
	return &v
}

func main() {
	p, err := personJSON("Gopher", time.Now().Add(time.Hour*24*365*-16))

	if err != nil {
		panic(err)
	}
	fmt.Println(string(p))
}
