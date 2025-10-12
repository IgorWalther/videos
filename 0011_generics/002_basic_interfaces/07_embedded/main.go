package main

import "io"

type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var _ io.ReadWriter
	var _ io.ReadWriteCloser
}
