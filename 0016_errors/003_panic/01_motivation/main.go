package main

// 1. Unreachable code
// 2. Internal functions with strong recover invariant (e.g. parsing)

func GetCharByIndex(str string, idx int) rune {
	var counter int

	for _, r := range str {
		if counter == idx {
			return r
		}

		counter++
	}

	panic("unreachable")
}
