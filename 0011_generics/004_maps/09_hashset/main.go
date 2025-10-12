package main

import (
	"errors"
	"fmt"
)

var ErrKeyNotFound = errors.New("key not found")

type set[K comparable] struct {
	data map[K]any
}

func New[K comparable](capacity int) *set[K] {
	return &set[K]{
		data: make(map[K]any, capacity),
	}
}

func (s *set[K]) Set(key K, value any) {
	s.data[key] = value
}

func (s *set[K]) Get(key K) (any, error) {
	if v, ok := s.data[key]; ok {
		return v, nil
	} else { // del else?
		return v, ErrKeyNotFound
	}
}

func main() {
	s := New[int](10)
	s.Set(5, 5)
	fmt.Println(s.Get(5))
	fmt.Println(s.Get(6))
}
