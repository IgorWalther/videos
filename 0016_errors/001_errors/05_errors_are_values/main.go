package main

import (
	"bufio"
	"io"
)

func bufioExample(input io.Reader) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		token := scanner.Text()
		_ = token
		// handle
	}

	if err := scanner.Err(); err != nil {
		// check error
	}
}

// before
//	_, err = fd.Write(p0[a:b])
//	if err != nil {
//		return err
//	}
//	_, err = fd.Write(p1[c:d])
//	if err != nil {
//		return err
//	}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

func example(fd io.Writer, data []byte) error {
	ew := &errWriter{w: fd}
	ew.write(data[1:2])
	ew.write(data[3:4])
	ew.write(data[5:6])

	if ew.err != nil {
		return ew.err
	}

	return nil
}

// Errors are values

// Errors in Go are not exceptions; they are ordinary values that can be stored, passed around,
// and manipulated like any other value.

// Values can be programmed, and since errors are values, errors can be programmed

// You can write code that works with errors - not just check them.

// It’s critical that the program check the errors however they are exposed.

// The goal isn’t to avoid checking errors, but to handle them gracefully.

// Use the language to simplify your error handling.

// Go gives you the tools to reduce repetitive if err != nil code — if you design your API wisely.

// Whatever you do, always check your errors!

// A reminder that no matter how elegant your abstraction, unchecked errors are still fatal.

// Errors are values and the full power of the Go programming language is available for processing them

// Treat errors as first-class citizens - part of the program’s data flow, not interruptions to it.
