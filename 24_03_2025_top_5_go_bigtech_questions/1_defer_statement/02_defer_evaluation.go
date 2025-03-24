package main

import "fmt"

func main() {
	deferEvaluationAsUsual()
	deferEvaluationOnDemand()
}

type data struct {
	value int
}

// The function value and parameters to the call are evaluated
// as usual and saved anew but the actual function is not invoked
// deferEvaluationAsUsual
func deferEvaluationAsUsual() {
	d := data{}
	defer fmt.Println(d.value) // 5

	d.value++
}

func deferEvaluationOnDemand() {
	d := data{}
	defer func() {
		fmt.Println(d.value) // 6
	}()

	d.value++
}
