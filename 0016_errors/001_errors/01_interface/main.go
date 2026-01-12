package main

type error interface {
	Error() string
}

// errors are values !!!

// It is the conventional interface for representing an error condition,
// with the nil value representing no error
