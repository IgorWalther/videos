package main

// T - type parameter
// int - type argument

func isEqual[T comparable](lhs, rhs T) bool {
	return lhs == rhs
}

func Less[T Ordered](x, y T) bool {
	return (isNaN(x) && !isNaN(y)) || x < y
}

func isNaN[T Ordered](x T) bool {
	return x != x
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
