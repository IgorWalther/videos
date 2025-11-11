package matrix

import (
	"math/rand"
	"runtime"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func makeRandMat(r, c int) *matrix {
	m := New(r, c)
	rnd := rand.New(rand.NewSource(42))

	for i := range m.Data {
		m.Data[i] = rnd.Float64()
	}

	return m
}

const (
	C = 1024
	R = 4096
)

func TestCorrectness(t *testing.T) {
	A := makeRandMat(R, C)
	B := makeRandMat(R, C)

	seq, err := AddSeq(A, B)
	require.NoError(t, err)

	par, err := AddPar(A, B, runtime.NumCPU())
	require.NoError(t, err)

	require.Equal(t, seq, par)
}

func BenchmarkAdd(b *testing.B) {
	b.Run("seq", func(b *testing.B) {
		b.ReportAllocs()

		A := makeRandMat(R, C)
		B := makeRandMat(R, C)

		for b.Loop() {
			C, _ := AddSeq(A, B)
			_ = C.Data[len(C.Data)-1]
		}
	})

	// check runtime.NumCPU() and runtime.NumCPU() * 100

	// why?
	// BenchmarkAdd/par/w=120-12            837           1400609 ns/op        33565177 B/op        243 allocs/op
	// BenchmarkAdd/par/w=1200-12           642           1873310 ns/op        33660769 B/op       2404 allocs/op
	for _, workers := range []int{1, 2, 4, 8, runtime.NumCPU(), runtime.NumCPU() * 3, runtime.NumCPU() * 10,
		runtime.NumCPU() * 100} {
		b.Run("par/w="+strconv.Itoa(workers), func(b *testing.B) {
			b.ReportAllocs()

			A := makeRandMat(R, C)
			B := makeRandMat(R, C)

			for b.Loop() {
				C, _ := AddPar(A, B, workers)
				_ = C.Data[len(C.Data)-1]
			}
		})
	}
}
