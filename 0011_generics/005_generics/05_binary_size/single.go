//go:build single

package main

type Number interface {
	~int
}

type Number2 interface {
	~int
}

var Sink any

//go:noinline
func keep(v any) {
	Sink = v
}

func Sum[T Number, S Number2](lhs, rhs T, a S) T {
	return lhs + rhs
}

// go build -tags=single -o single -gcflags='-N -l' single.go
func main() {
	keep(Sum[int](1, 2, 3))
}
