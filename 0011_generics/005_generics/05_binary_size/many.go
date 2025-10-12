//go:build many

package main

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Number2 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Sum[T Number, S Number2](lhs, rhs T, a S) T {
	return lhs + rhs
}

var Sink any

//go:noinline
func keep(v any) {
	Sink = v
}

// go build -tags=many -o many -gcflags='-N -l' many.go
func main() {
	keep(Sum[int](1, 2, 3))
	keep(Sum[uint8](1, 2, 3))
	keep(Sum[int16](1, 2, 3))
	keep(Sum[int32](1, 2, 3))
	keep(Sum[int64](1, 2, 3))
	keep(Sum[uintptr](1, 2, 3))
	keep(Sum[float64](1, 2, 3))
	keep(Sum[float32](1, 2, 3))
	keep(Sum[int32](1, 2, 3))
	keep(Sum[int64](1, 2, 3))
	keep(Sum[uint](1, 2, 3))
}

//go tool nm many | grep 'Sum\['
//
//10009f4d0 R main..dict.Sum[float32,int]
//10009f4c0 R main..dict.Sum[float64,int]
//10009f460 R main..dict.Sum[int,int]
//10009f480 R main..dict.Sum[int16,int]
//10009f490 R main..dict.Sum[int32,int]
//10009f4a0 R main..dict.Sum[int64,int]
//10009f4e0 R main..dict.Sum[uint,int]
//10009f470 R main..dict.Sum[uint8,int]
//10009f4b0 R main..dict.Sum[uintptr,int]
//100069ae0 T main.Sum[go.shape.float32,go.shape.int]
//100069b20 T main.Sum[go.shape.float64,go.shape.int]
//100069ca0 T main.Sum[go.shape.int,go.shape.int]
//100069c20 T main.Sum[go.shape.int16,go.shape.int]
//100069be0 T main.Sum[go.shape.int32,go.shape.int]
//100069ba0 T main.Sum[go.shape.int64,go.shape.int]
//100069aa0 T main.Sum[go.shape.uint,go.shape.int]
//100069c60 T main.Sum[go.shape.uint8,go.shape.int]
//100069b60 T main.Sum[go.shape.uintptr,go.shape.int]
