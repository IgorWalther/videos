package main

import (
	"crypto/rand"
	"testing"
	"testing/cryptotest"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	cryptotest.SetGlobalRandom(t, 42)

	p1, _ := rand.Prime(nil, 32)
	p2, _ := rand.Prime(nil, 32)
	p3, _ := rand.Prime(nil, 32)

	got := [3]int64{p1.Int64(), p2.Int64(), p3.Int64()}
	want := [3]int64{3713413729, 3540452603, 4293217813}

	require.Equal(t, want, got)
}
