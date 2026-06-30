package _1_list

import "sync"

type Node struct {
	value int
	next  *Node
	mx    *sync.Mutex
}

type FineGrainedList struct {
	head *Node
	// mx sync.Mutex
}

//func (l *FineGrainedList) Insert(val int) {
//	newNode := &Node{value: val}
//	current := l.head
//
//	for current != nil {
//		current.mx.Lock()
//		if current.next == nil {
//			current.next = newNode
//			current.mx.Unlock()
//			return
//		}
//		next := current.next
//		current.mx.Unlock()
//		current = next
//	}
//}

func (l *FineGrainedList) Insert(val int) {
	newNode := &Node{value: val}
	current := l.head

	for current != nil {
		withLock(current.mx, func() {
			if current.next == nil {
				current.next = newNode
				return
			}

			next := current.next
			current.mx.Unlock()
			current = next
		})
	}
}

func withLock(l sync.Locker, action func()) {
	if action == nil {
		return
	}

	l.Lock()
	defer l.Unlock()

	action()
}
