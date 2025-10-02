package main

const (
	someConst    string = "123"
	anotherConst        = "123"
)

type MyTypeStr string

const login MyTypeStr = "123"

var name1 = "123"

func consumer(a MyTypeStr) {
	name1 = "123"
}

const (
	_ = 1 << (10 * iota)
	Kb
	Mb
	Gb
	Tb
)
