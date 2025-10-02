package main

import (
	"fmt"
	"math/bits"
)

// An integer, floating-point, or complex type represents the set of integer,
// floating-point, or complex values, respectively.
// They are collectively called numeric types. The predeclared architecture-independent numeric types are:
func main() {
	var _ uint8  // the set of all unsigned  8-bit integers (0 to 255)
	var _ uint16 // the set of all unsigned 16-bit integers (0 to 65535)
	var _ uint32 // the set of all unsigned 32-bit integers (0 to 4294967295)
	var _ uint64 // the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	var _ int8  // the set of all signed  8-bit integers (-128 to 127)
	var _ int16 // the set of all signed 16-bit integers (-32768 to 32767)
	var _ int32 // the set of all signed 32-bit integers (-2147483648 to 2147483647)
	var _ int64 // the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	var _ float32 // the set of all IEEE 754 32-bit floating-point numbers
	var _ float64 // the set of all IEEE 754 64-bit floating-point numbers

	var _ complex64  // the set of all complex numbers with float32 real and imaginary parts
	var _ complex128 // the set of all complex numbers with float64 real and imaginary parts

	// uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
	var _ uintptr

	var _ byte // alias for uint8
	var _ rune // alias for int32

	// builtin.go

	// There is also a set of predeclared integer types with implementation-specific sizes:
	// uint     either 32 or 64 bits
	// int      same size as uint
	fmt.Println(bits.UintSize)
}
