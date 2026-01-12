package main

import (
	"errors"
	"fmt"
	"io"
)

func CatchT[T any](handler func(T)) func(any) bool {
	return func(e any) bool {
		if v, ok := e.(T); ok {
			handler(v)
			return true
		}

		return false
	}
}

func ChainCatch(fn func(), finally func(), catchers ...func(any) bool) {
	defer func() {
		if finally != nil {
			defer finally()
		}

		if r := recover(); r != nil {
			defer func() {
				if rr := recover(); rr != nil {
					panic(rr)
				}
			}()

			for _, c := range catchers {
				if c(r) {
					return
				}
			}

			panic(r)
		}
	}()

	fn()
}

func main() {
	ChainCatch(
		func() {
			fmt.Println("start")
			//panic(io.EOF)
			panic("string")
			//panic(errors.New("some error"))
		},
		func() {
			fmt.Println("finally")
		},
		CatchT[error](func(err error) {
			if errors.Is(err, io.EOF) {
				fmt.Println("caught io.EOF specifically")
				return
			}

			fmt.Println("caught general error:", err)
		}),
		CatchT[string](func(s string) {
			fmt.Println("caught string:", s)
		}),
	)
}
