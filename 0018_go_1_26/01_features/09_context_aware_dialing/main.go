package main

import (
	"context"
	"fmt"
	"net"
	"net/netip"
)

func main() {
	var d net.Dialer
	ctx := context.Background()
	raddr := netip.MustParseAddrPort("127.0.0.1:8080")
	conn, err := d.DialTCP(ctx, "tcp", netip.AddrPort{}, raddr)
	fmt.Printf("connected, err=%v\n", err)
	defer conn.Close()

	// d.DialUDP()
	// d.DialIP()
}
