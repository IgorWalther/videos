package main

import "fmt"

// For readability, an underscore character _ may appear after a base prefix or between successive digits;
// such underscores do not change the literal's value.

func main() {
	fmt.Println(42)
	fmt.Println(4_2)
	fmt.Println(0600)
	fmt.Println(0_600)
	fmt.Println(0o600)
	fmt.Println(0o600)
	fmt.Println(0xBad)
	fmt.Println(0xB_a_d)
	fmt.Println(0x_67_7a_2f_cc_40_c6)
	fmt.Println(170141183460)
	fmt.Println(170_141183_460)
	fmt.Println(0_0_0_0_0_0_0_0_0)

	//_42         // an identifier, not an integer literal
	//42_         // invalid: _ must separate successive digits
	//4__2        // invalid: only one _ at a time
	//0_xBadFace  // invalid: _ must separate successive digits
}
