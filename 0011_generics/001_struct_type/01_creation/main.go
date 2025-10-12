package main

type T struct {
	a int
	b int
	c string
}

func main() {
	_ = new(T)

	_ = &T{
		a: 1,
	}

	_ = T{
		b: 2,
	}
}
