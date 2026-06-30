// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"sync/atomic"
)

// sync.WaitGroup{}

// see sync.runtime_SemacquireWaitGroup, go linkname restricted
func runtime_SemacquireWaitGroup(addr *uint32) {}

// see sync.runtime_Semrelease, go linkname restricted
func runtime_Semrelease(s *uint32, handoff bool, skipframes int) {}

type WaitGroupAdvanced struct {
	state atomic.Uint64
	sema  uint32
}

func (wg *WaitGroupAdvanced) Add(delta int) {
	// state
	// v: 0000000001100000000110000110001 w: 1110101010101010101001010101010101

	state := wg.state.Add(uint64(delta) << 32)
	v := int32(state >> 32)
	w := uint32(state)

	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}

	if w != 0 && delta > 0 && v == int32(delta) {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}

	if v > 0 || w == 0 {
		return
	}

	if wg.state.Load() != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}

	wg.state.Store(0)
	for ; w != 0; w-- {
		runtime_Semrelease(&wg.sema, false, 0)
	}
}

// Done decrements the [WaitGroup] counter by one.
func (wg *WaitGroupAdvanced) Done() {
	wg.Add(-1)
}

// Wait blocks until the [WaitGroup] counter is zero.
func (wg *WaitGroupAdvanced) Wait() {
	for {
		state := wg.state.Load()
		v := int32(state >> 32)

		if v == 0 {
			return
		}

		if wg.state.CompareAndSwap(state, state+1) {
			runtime_SemacquireWaitGroup(&wg.sema)

			if wg.state.Load() != 0 {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}

			return
		}
	}
}
