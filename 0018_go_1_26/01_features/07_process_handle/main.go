package main

import (
	"fmt"
	"os"
)

func main() {
	attr := &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}
	proc, _ := os.StartProcess("/bin/echo", []string{"echo", "hello, @igoroutine"}, attr)
	defer proc.Wait()

	fmt.Println("pid =", proc.Pid)

	err := proc.WithHandle(func(handle uintptr) {
		fmt.Println("handle =", handle)
	})

	fmt.Println(err)
}
