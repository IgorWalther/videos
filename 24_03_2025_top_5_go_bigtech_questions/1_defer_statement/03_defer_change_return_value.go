package main

import "fmt"

func main() {
	fmt.Println(deferChangeReturnValue().value)
}

type returnData struct {
	value int
}

// The function value and parameters to the call are evaluated
// as usual and saved anew but the actual function is not invoked
// deferEvaluationAsUsual
func deferChangeReturnValue() (d returnData) {
	defer func() {
		d.value = 43
	}()

	d.value = 42
	return
}
