package readallbench

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

// ReadAll now allocates less intermediate memory and returns a minimally sized final slice.
// It is often about two times faster while typically allocating around half as much total memory,
// with more benefit for larger inputs.

var (
	keepBytes []byte
	keepCap   int
)

func measureReadAll(b *testing.B, n int) {
	src := bytes.Repeat([]byte("a"), n)
	b.ReportAllocs()

	for b.Loop() {
		r := bytes.NewReader(src)
		out, err := io.ReadAll(r)
		if err != nil {
			b.Fatalf("ReadAll: %v", err)
		}

		keepBytes = out
		keepCap = cap(out)

		require.Equal(b, len(out), n)
	}
}

func BenchmarkReadAll_1KiB(b *testing.B) {
	measureReadAll(b, 1<<10)
}

func BenchmarkReadAll_4KiB(b *testing.B) {
	measureReadAll(b, 4<<10)
}

func BenchmarkReadAll_16KiB(b *testing.B) {
	measureReadAll(b, 16<<10)
}

func BenchmarkReadAll_65KiB(b *testing.B) {
	measureReadAll(b, 65<<10)
}

func BenchmarkReadAll_256KiB(b *testing.B) {
	measureReadAll(b, 256<<10)
}

func BenchmarkReadAll_1MiB(b *testing.B) {
	measureReadAll(b, 1<<20)
}

func BenchmarkReadAll_4MiB(b *testing.B) {
	measureReadAll(b, 4<<20)
}
