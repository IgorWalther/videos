package mutexbench

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var _ sync.Locker = (*ticketLock)(nil)

func NewTicketLock() *ticketLock {
	return &ticketLock{}
}

type ticketLock struct {
	ownerTicket    atomic.Int64
	nextFreeTicket atomic.Int64
}

func (t *ticketLock) Lock() {
	ticket := t.nextFreeTicket.Add(1) - 1

	for t.ownerTicket.Load() != ticket {
		runtime.Gosched()
	}
}

func (t *ticketLock) Unlock() {
	t.ownerTicket.Add(1)
}
