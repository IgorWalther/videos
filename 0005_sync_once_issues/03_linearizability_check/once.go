package linearizability

import (
	"sync/atomic"
)

type myOnce struct {
	done atomic.Uint32
}

func (o *myOnce) Do(f func()) {
	if o.done.CompareAndSwap(0, 1) {
		f()
	}
}
