package main

import (
	"sync"
)

func main() {
	r := Reentrant{
		mx: new(sync.Mutex),
	}

	r.Outer()
}

type Reentrant struct {
	mx *sync.Mutex
}

//func (r *Reentrant) Push() int {
//	r.mx.Lock()
//	defer r.mx.Unlock()
//
//	const capacity = 10
//	if r.Size() < capacity {
//		// ...
//		panic(nil)
//	}
//}
//
//func (r *Reentrant) Size() int {
//	r.mx.Lock()
//	defer r.mx.Unlock()
//
//	return 1
//}

func (r *Reentrant) Outer() {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.Inner()
}

// for example, tree

func (r *Reentrant) Inner() {
	r.mx.Lock()
	defer r.mx.Unlock()
}
