package main

import (
	"fmt"
	"net/netip"
	"slices"
)

func main() {
	prefixes := []netip.Prefix{
		// ...
	}

	// Prefixes sort first by validity (invalid before valid), then
	// address family (IPv4 before IPv6), then masked prefix address, then
	// prefix length, then unmasked address.

	slices.SortFunc(prefixes, netip.Prefix.Compare)

	for _, p := range prefixes {
		fmt.Println(p.String())
	}
}
