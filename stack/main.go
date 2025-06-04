package main

type Node[T any] struct {
	value T
	prev  *Node[T]
}

type Stack[T any] struct {
	size int
	head *Node[T]
}

func (s *Stack[T]) Push(item T) {
	node := &Node[T]{
		value: item,
		prev:  s.head,
	}
	s.size++
	s.head = node
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}
	s.size--

	head := s.head
	s.head = s.head.prev
	head.prev = nil
	return head.value, true
}

func (s Stack[T]) Peek() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}

	return s.head.value, true
}
