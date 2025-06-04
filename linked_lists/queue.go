package main

type Node[T any] struct {
	value T
	Next  *Node[T]
}

type Queue[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) Push(value T) {
	var node = &Node[T]{
		value: value,
		Next:  nil,
	}
	q.size++
	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.Next = node
	q.tail = node
}

func (q *Queue[T]) Pop() *T {
	if q.size == 0 {
		return nil
	}
	q.size--
	head := q.head
	q.head = q.head.Next
	head.Next = nil

	return &head.value
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}
	return q.head.value, true
}

func main() {}
