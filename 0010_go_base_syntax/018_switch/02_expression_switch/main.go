package main

import (
	"fmt"
	"os"
)

func main() {
	var day string
	fmt.Fscan(os.Stdin, &day)

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Wednesday":
		fmt.Println("Middle of the week")
	case "Friday":
		fmt.Println("Almost weekend")
	case "Saturday", "Sunday": // !
		fmt.Println("Weekend")
	default:
		fmt.Println("Some other day")
	}
}
