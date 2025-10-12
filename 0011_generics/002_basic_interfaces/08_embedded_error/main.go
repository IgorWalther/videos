package main

type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

// When embedding interfaces, methods with the same names must have identical signatures.
//type ReadCloser interface {
//	Reader   // includes methods of Reader in ReadCloser's method set
//	Close()  // illegal: signatures of Reader.Close and Close are different
//}
