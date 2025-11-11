package main

import "sync/atomic"

type Counter struct {
	count uint32
}

func (p *Counter) SetCount(n uint32) {
	atomic.StoreUint32(&p.count, n)
}

func (p *Counter) Count() uint32 {
	return atomic.LoadUint32(&p.count)
}
