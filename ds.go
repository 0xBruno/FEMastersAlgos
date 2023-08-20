package main

import (
	"fmt"
)

type Node struct {
	data any
	next *Node
	prev *Node
}

// LinkedListInterface API
type LinkedListInterface interface {
	getLength() int
	insertAt(item any, index int) error
	append(item any)
	//remove(item any) error
	//removeAt(index int) error
	//prepend(item any)
	//get(index int) error
}

// LinkedList implementation
type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (l *LinkedList) getLength() int {
	return l.length
}

// insertAt inserts any data at any index within the list
func (l *LinkedList) insertAt(item any, index int) error {
	node := l.head
	outOfBoundsError := fmt.Errorf("cannot insert at an out of bounds index into the list or list is empty. Use append() instead")

	if index > l.length {
		return outOfBoundsError
	} else if l.length == 0 {
		return outOfBoundsError
	}

	for i := 0; ; i++ {
		if index == i {
			node.data = item
			return nil
		}

		node = node.next
	}
}

func (l *LinkedList) append(item any) {

	node := Node{
		data: item,
		next: nil,
		prev: nil,
	}

	// empty list
	if l.length == 0 {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	current := l.head
	for {
		// if we are at the end
		if current.next == nil {
			current.next = &node
			l.tail = &node
			l.length++
			return
		}

		current = current.next
	}
}

func NewLinkedList() LinkedList {
	return LinkedList{
		length: 0,
		head: &Node{
			data: nil,
			next: nil,
			prev: nil,
		},
	}
}
