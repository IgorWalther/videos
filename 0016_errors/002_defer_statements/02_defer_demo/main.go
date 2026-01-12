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

	defer in.Close()

	out, err := os.Create(dst)

	if err != nil {
		return err
	}

	defer out.Close()

	const bufferSize = 1 << 10
	buf := make([]byte, bufferSize)

	for {
		n, err := in.Read(buf)

		if n <= 0 || err != nil {
			break
		}

		if _, wErr := out.Write(buf[:n]); wErr != nil {
			return wErr
		}
	}

	return nil
}

const (
	srcFile = "src.txt"
	dstFile = "dst.txt"
)

// go run main.go

// 1. Обработка исключительных ситуаций приложения (panic middleware)
// 2. Закрытие ресурсов (aka finally block)

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
