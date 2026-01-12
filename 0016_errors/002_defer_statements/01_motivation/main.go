package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)

	if err != nil {
		return err
	}

	out, err := os.Create(dst)

	if err != nil {
		in.Close()
		return err
	}

	const bufferSize = 1 << 10
	buf := make([]byte, bufferSize)

	for {
		n, err := in.Read(buf)

		if n <= 0 || err != nil {
			break
		}

		if _, wErr := out.Write(buf[:n]); wErr != nil {
			in.Close()
			out.Close()
			return wErr
		}
	}

	in.Close()
	out.Close()

	return nil
}

const (
	srcFile = "src.txt"
	dstFile = "dst.txt"
)

// go run main.go

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Println("can not get working directory: " + err.Error())
		return
	}

	src := filepath.Join(wd, srcFile)
	dst := filepath.Join(wd, dstFile)

	if err := copyFile(src, dst); err != nil {
		fmt.Println("error:", err)
	}
}
