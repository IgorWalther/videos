//go:build performance_test

package tests

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerformance(t *testing.T) {
	r := testing.Benchmark(func(b *testing.B) {
		for b.Loop() {
			rand.N[int](10e9)
		}
	})

	require.Zero(t, r.AllocsPerOp())
}
