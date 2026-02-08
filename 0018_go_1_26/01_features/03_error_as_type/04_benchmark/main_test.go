package main

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

type AppError struct {
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	appErrChain = func() error {
		return fmt.Errorf("wrap1: %w",
			fmt.Errorf("wrap2: %w",
				&AppError{
					Message: "database is down",
				},
			),
		)
	}()

	netErrChain = func() error {
		base := &net.DNSError{
			Name: "example.com",
			Err:  "no such host",
		}

		return fmt.Errorf("wrap1: %w", fmt.Errorf("wrap2: %w", base))
	}()
)

func BenchmarkAs(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		var target *AppError
		_ = errors.As(appErrChain, &target)
	}
}

func BenchmarkAsType(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		_, _ = errors.AsType[*AppError](appErrChain)
	}
}

func BenchmarkAsMulti(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		var opErr *net.OpError
		if errors.As(netErrChain, &opErr) {
			continue
		}

		var dnsErr *net.DNSError
		_ = errors.As(netErrChain, &dnsErr)
	}
}

func BenchmarkAsType_Multi(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		if _, ok := errors.AsType[*net.OpError](netErrChain); ok {
			continue
		}

		_, _ = errors.AsType[*net.DNSError](netErrChain)
	}
}
