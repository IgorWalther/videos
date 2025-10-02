package main

func main() {
	s := make([]int64, 10)
	s[100] = 1 // panic!
}

func foo(v int) {
	if v < 10 {
		panic("invariant error") // !!!
	}
}
