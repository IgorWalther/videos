package main

// F-bounded quantification or recursively bounded quantification

type Builder[A Builder[A]] interface {
	WithTimeout(ms int) A
	WithDelay(ms int) A
}

type Bad struct{}

func (b Bad) WithTimeout(ms int) int {
	return 123
}

func (b Bad) WithDelay(ms int) int {
	return 123
}

type Good struct {
	timeout int
	delay   int
}

func (b Good) WithTimeout(ms int) Good {
	b.timeout = ms
	return b
}

func (b Good) WithDelay(ms int) Good {
	b.delay = ms
	return b
}

func main() {
	g := Good{}.
		WithTimeout(100).
		WithDelay(100)
	_ = g

	//b := Bad{}.
	//	WithTimeout(100). // ???

	// Standard(Bad{})
}

func Standard[A Builder[A]](b A) A {
	return b.WithTimeout(1000).
		WithDelay(3)
}
