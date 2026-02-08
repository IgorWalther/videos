//go:build unix

package perfbench

import (
	"syscall"
	"testing"
)

func BenchmarkSyscallGetpid(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		_, _, _ = syscall.RawSyscall(syscall.SYS_GETPID, 0, 0, 0)
	}
}

func BenchmarkSyscallGetpidParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, _ = syscall.RawSyscall(syscall.SYS_GETPID, 0, 0, 0)
		}
	})
}
