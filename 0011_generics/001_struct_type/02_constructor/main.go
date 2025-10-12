package main

type T struct {
	a int
	b int
	c string
}

// func NewT(a, b int, c string) *t for private

func NewT(a, b int, c string) *T {
	return &T{
		a: a,
		b: b,
		c: c,
	}
}
