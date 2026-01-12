package main

import (
	"errors"
	"io"
	"os"
)

func copyAndClose(dst io.WriteCloser, src io.Reader) (err error) {
	_, err = io.Copy(dst, src)
	closeErr := dst.Close()

	if err != nil && closeErr != nil {
		return errors.Join(err, closeErr)
	}

	if err != nil {
		return err
	}

	return closeErr
}

// if errors.Is(err, io.EOF) { /* ... */ }

func do() (err error) {
	f, err := os.Open("file")
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			err = errors.Join(err, cerr) // !!!
		}
	}()

	// ...
	return nil
}
