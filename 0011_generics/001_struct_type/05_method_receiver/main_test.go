package main

import "testing"

func BenchmarkReceiver(b *testing.B) {
	b.Run("pointer receiver", func(b *testing.B) {
		b.ReportAllocs()

		t := T1{}

		for b.Loop() {
			t.Foo()
		}
	})

	b.Run("value receiver", func(b *testing.B) {
		b.ReportAllocs()

		t := T2{}

		for b.Loop() {
			t.Foo()
		}
	})
}

type T1 struct {
	_ [1024]byte
}

type T2 struct {
	_ [1024]byte
}

//go:noinline
func (t *T1) Foo() {}

//go:noinline
func (t T2) Foo() {}
