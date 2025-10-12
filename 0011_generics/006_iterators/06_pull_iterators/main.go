package main

import (
	"fmt"
	"iter"
)

func Range(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	seq := Range(5)

	next, stop := iter.Pull(seq)
	defer stop()

	for {
		v, ok := next()

		if !ok {
			break
		}

		fmt.Println(v)
	}
}
