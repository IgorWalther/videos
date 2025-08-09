package main

// before Go 1.25:
// panic: PANIC [recovered]
//   panic: PANIC
// after Go 1.25:
// panic: PANIC [recovered, repanicked]

func main() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	panic("PANIC")
}
