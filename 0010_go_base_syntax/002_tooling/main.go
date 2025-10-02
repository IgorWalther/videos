package main

import "fmt"

// GOPATH
// GOROOT

//import (
//	"flag"
//	"log"
//
//	firstApi "myapp/src/apiserver"
//	secondApi "myapp/sørc/bpiserver"
//
//	"database/sql"
//	_ "github.com/mattn/go-sqlite3"
//)

// go install golang.org/x/tools/cmd/goimports@latest

//Разбиваем импорты на группы
// * стандартная библиотека
// * модули проекта
// * сторонние библиотеки
// внутри каждой группы сортируем
// Циклические импорты = CE
// Можно именовать в camelCase
// Можно именовать wildcard

// go mod init (todo: workspace)
// go build
// go run main.go
// go test

const name = "World"

func main() {
	fmt.Println(greater(name))
}

func greater(name string) string {
	return "Hello, " + name + "!"
}
