package main

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	value T
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		nil,
		value,
	}
}

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.length++

	node := NewNode(item)

	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}
func (q *Queue[T]) Deque() interface{} {
	if q.length == 0 {
		return nil
	}
	q.length--

	head := q.head
	q.head = q.head.next
	return head.value

}

func (q *Queue[T]) Peek() interface{} {
	return q.head.value
}

func main() {
	q := NewQueue[string]()

	q.Enqueue("first value")
	fmt.Println(q.Peek())
	q.Enqueue("second value")
	fmt.Println(q.Peek())
	q.Enqueue("third value")

	fmt.Println(q.Peek())
	item := q.Deque()
	fmt.Println(q.Peek())
	fmt.Print(item)

}
