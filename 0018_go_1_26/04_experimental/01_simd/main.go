//go:build goexperiment.simd

package main

import "simd/archsimd"

func main() {
	_ = archsimd.X86
	// see SIMD video and report
}
