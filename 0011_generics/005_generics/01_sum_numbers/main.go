package main

type Number interface {
	int64 | float64 | float32 | int32 | int
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}
