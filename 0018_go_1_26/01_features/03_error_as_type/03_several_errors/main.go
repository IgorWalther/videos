package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	var err error = &net.DNSError{
		Name: "example.com",
		Err:  "no such host",
	}

	if opErr, ok := errors.AsType[*net.OpError](err); ok {
		fmt.Println("Network op failed:", opErr.Op)
	} else if dnsErr, ok := errors.AsType[*net.DNSError](err); ok {
		fmt.Println("DNS failed:", dnsErr.Name, dnsErr.Err)
	} else {
		fmt.Println("Unknown error:", err)
	}
}
