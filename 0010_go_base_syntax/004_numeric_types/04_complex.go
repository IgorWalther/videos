package main

import "fmt"

func main() {
	fmt.Println(0i)
	fmt.Println(0o123i) // == 0o123 * 1i == 83i
	fmt.Println(0xabci) // == 0xabc * 1i == 2748i
	fmt.Println(0.i)
	fmt.Println(2.71828i)
	fmt.Println(1.e+0i)
	fmt.Println(6.67428e-11i)
	fmt.Println(1e6i)
	fmt.Println(.25i)
	fmt.Println(.12345e+5i)
	fmt.Println(0x1p-2i) // == 0x1p-2 * 1i == 0.25i

	a := complex(1, 2.5)
	fmt.Println(real(a))
	fmt.Println(imag(a))
}
