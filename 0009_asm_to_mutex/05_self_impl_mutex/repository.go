package mutexbench

import (
	"math/rand/v2"
	"sync"
	"time"
)

func NewRepository(mx sync.Locker) *repository {
	return &repository{
		mx: mx,
		m:  make(map[int]string),
	}
}

type repository struct {
	mx sync.Locker
	m  map[int]string
}

//go:noinline
func (r *repository) Operation() {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.m[rand.N[int](10e9)] = "hello"
	time.Sleep(time.Millisecond)
}
