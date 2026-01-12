package main

import (
	"io"
	"testing"
)

var (
	catchError  = CatchT[error](func(err error) {})
	catchString = CatchT[string](func(s string) {})
)

const computeIterations = 16

var sink int

func doCompute(n int) {
	s := 0

	for i := 0; i < n; i++ {
		s += i
	}

	sink = s
}

func doErr(fail bool) error {
	if fail {
		return io.EOF
	}

	doCompute(computeIterations)
	return nil
}

func BenchmarkErrorSuccess(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if err := doErr(false); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkErrorFailure(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if err := doErr(true); err != nil {
			_ = err
		}
	}
}

func BenchmarkTryCatchSuccess(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ChainCatch(
			func() {
				doCompute(computeIterations)
			},
			nil,
			catchError, catchString,
		)
	}
}

func BenchmarkTryCatchFailure(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ChainCatch(
			func() {
				panic(io.EOF)
			},
			nil,
			catchError, catchString,
		)
	}
}
