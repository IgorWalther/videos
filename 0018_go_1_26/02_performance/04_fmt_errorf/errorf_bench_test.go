package errorfbench

import (
	"fmt"
	"testing"
)

var keep error

func BenchmarkErrors(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		keep = fmt.Errorf("foo")
	}
}
