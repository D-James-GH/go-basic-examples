package main

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	value T
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func (L *LinkedList[T]) Insert(value T) {
	node := &Node[T]{
		value: value,
	}
	L.len++
	if L.head == nil {
		L.head = node
		L.tail = node
	} else {
		L.tail.next = node
		node.prev = L.tail
		L.tail = node
	}
}
func (L *LinkedList[T]) RemoveLast() *Node[T] {
	L.len--
	if L.tail != nil {
		node := L.tail
		L.tail = node.prev
		L.tail.next = nil
		node.prev = nil
		return node
	}
	return nil
}

func (L *LinkedList[T]) RemoveFirst() *Node[T] {
	L.len--
	if L.head != nil {
		newHead := L.head.next
		L.head.next = nil
		newHead.prev = nil
		L.head = newHead
		return newHead
	}
	return nil
}

func (L *LinkedList[T]) GetAt(i int) *Node[T] {
	if i >= L.len {
		return nil
	}
	if i < L.len/2 {
		// go from head
		index := 0
		node := L.head
		for index < i {
			node = node.next
			index++
		}
		return node
	} else {
		// go from tail
		i = (L.len) - i - 1
		index := 0
		node := L.tail
		for index < i {
			node = node.prev
			index++
		}
		return node
	}
}
func (L *LinkedList[T]) RemoveAt(i int) *Node[T] {
	if i >= L.len {
		return nil
	}
	var node *Node[T]
	if i == 0 {
		node = L.RemoveFirst()
	} else if i == L.len-1 {
		node = L.RemoveLast()
	} else {
		node = L.GetAt(i)
		// last node
		prev := node.prev
		next := node.next
		// link them
		prev.next = next
		next.prev = prev

		node.next = nil
		node.prev = nil
	}

	L.len--
	return node
}

// print the linked list
func (L *LinkedList[T]) String() string {
	var s string
	node := L.head

	for node != nil {
		s += fmt.Sprintf("%v", node.value)
		node = node.next
	}
	return s
}

func main() {
	list := LinkedList[int]{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	s := list.String()
	fmt.Println(s)
	// list.RemoveLast()
	// list.RemoveFirst()
	list.RemoveAt(5)
	s = list.String()
	fmt.Println(s)

	list.RemoveAt(0)
	s = list.String()
	fmt.Println(s)

}
