package main

// see atomic_xxx.go

// Package atomic provides low-level atomic memory primitives
// useful for implementing synchronization algorithms.

// These functions require great care to be used correctly.

// These functions require great care to be used correctly.
// Except for special, low-level applications, synchronization is better
// done with channels or the facilities of the [sync] package.

// BUG(rsc): On 386, the 64-bit functions use instructions unavailable before the Pentium MMX.
// On non-Linux ARM, the 64-bit functions use instructions unavailable before the ARMv6k core.

// The booleans in ARM64 contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
//var ARM64 struct {
//	_          CacheLinePad
//	HasAES     bool
//	HasPMULL   bool
//	HasSHA1    bool
//	HasSHA2    bool
//	HasSHA512  bool
//	HasCRC32   bool
//	HasATOMICS bool
//	HasCPUID   bool
//	HasDIT     bool
//	IsNeoverse bool
//	_          CacheLinePad
//}
