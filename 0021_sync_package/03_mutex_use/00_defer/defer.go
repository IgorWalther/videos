package main

import "sync"

func main() {
	mx := new(sync.Mutex)

	mx.Lock()
	defer mx.Unlock()

	withLock(mx, func() {

	}) // default in Kotlin
}

// see sync.Locker
// foo
func withLock(mx sync.Locker, action func()) {
	mx.Lock()
	defer mx.Unlock()

	action()
}
