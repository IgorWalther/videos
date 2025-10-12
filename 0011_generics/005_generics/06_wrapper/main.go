package main

import (
	"container/list"
	"errors"
	"fmt"
)

var ErrListEmpty = errors.New("list is empty")

type linkedList[E any] struct {
	list *list.List
}

func New[E any]() *linkedList[E] {
	return &linkedList[E]{
		list: list.New(),
	}
}

func (l *linkedList[E]) PushBack(v E) {
	l.list.PushBack(v)
}

func (l *linkedList[E]) PopBack() (E, error) {
	if b := l.list.Back(); b != nil {
		l.list.Remove(b)
		val := b.Value.(E)
		return val, nil
	}

	var zero E
	return zero, ErrListEmpty
}

// sync.Pool

func main() {
	l := New[int]()
	l.PushBack(10)
	fmt.Println(l.PopBack())
}
