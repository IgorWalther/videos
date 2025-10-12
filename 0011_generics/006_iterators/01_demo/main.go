package main

type Set[E comparable] struct {
	m map[E]struct{}
}

func New[E comparable]() *Set[E] {
	return &Set[E]{m: make(map[E]struct{})}
}

func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

func (s *Set[E]) Contains(v E) bool {
	_, ok := s.m[v]
	return ok
}

func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
	r := New[E]()

	// for/range over internal Set field m.
	for v := range s1.m {
		r.Add(v)
	}

	for v := range s2.m {
		r.Add(v)
	}

	return r
}
