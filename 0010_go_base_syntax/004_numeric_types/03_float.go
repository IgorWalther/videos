package main

import "fmt"

// IEEE 754-2008 §5.12.3
func main() {
	fmt.Println(0.)
	fmt.Println(72.40)
	fmt.Println(0072.40) // == 72.40
	fmt.Println(2.71828)
	fmt.Println(1.e+0)
	fmt.Println(6.67428e-11)
	fmt.Println(1e6)
	fmt.Println(.25)
	fmt.Println(.12345e+5)
	fmt.Println(1_5.)      // == 15.0
	fmt.Println(0.15e+0_2) // == 15.0
	fmt.Println(0.15e+0_2) // == 15.0

	fmt.Println(0x1p-2)      // == 0.25
	fmt.Println(0x2.p10)     // == 2048.0
	fmt.Println(0x1.Fp+0)    // == 1.9375
	fmt.Println(0x.8p-0)     // == 0.5
	fmt.Println(0x_1FFFp-16) // == 0.1249847412109375
	fmt.Println(0x15e - 2)   // == 0x15e - 2 (integer subtraction)

	//0x.p1        // invalid: mantissa has no digits
	//1p-2         // invalid: p exponent requires hexadecimal mantissa
	//0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
	//1_.5         // invalid: _ must separate successive digits
	//1._5         // invalid: _ must separate successive digits
	//1.5_e1       // invalid: _ must separate successive digits
	//1.5e_1       // invalid: _ must separate successive digits
	//1.5e1_       // invalid: _ must separate successive digits
}
