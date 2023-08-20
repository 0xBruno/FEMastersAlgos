package main

import (
	"fmt"
)

type Node struct {
	Data any
	Next *Node
	Prev *Node
}

// LinkedListInterface API
type LinkedListInterface interface {
	GetLength() int
	InsertAt(item any, index int) error
	Append(item any)
	Get(index int) (any, error)
	// TODO
	// remove(item any) error
	//removeAt(index int) error
	//prepend(item any)
}

// LinkedList implementation
type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
	errs   map[string]error
}

func (l *LinkedList) GetLength() int {
	return l.Length
}

// InsertAt inserts any data at any index within the list
func (l *LinkedList) InsertAt(item any, index int) error {
	node := l.Head

	if index > l.Length {
		return l.errs["OutOfBoundInsertion"]
	} else if l.Length == 0 {
		return l.errs["OutOfBoundInsertion"]
	}

	for i := 0; ; i++ {
		if index == i {
			node.Data = item
			return nil
		}

		node = node.Next
	}
}

func (l *LinkedList) Append(item any) {

	node := Node{
		Data: item,
		Next: nil,
		Prev: nil,
	}

	// empty list
	if l.Length == 0 {
		l.Head = &node
		l.Tail = &node
		l.Length++
		return
	}

	current := l.Head
	for {
		// if we are at the end
		if current.Next == nil {
			current.Next = &node
			l.Tail = &node
			l.Length++
			return
		}

		current = current.Next
	}
}

func (l *LinkedList) Get(index int) (any, error) {
	current := l.Head
	for i := 0; i < l.Length-1; i++ {
		if index == i {
			return current.Data, nil
		}
		current = current.Next
	}
	return nil, l.errs["OutOfBoundGeneric"]
}

func NewLinkedList() LinkedList {
	return LinkedList{
		Length: 0,
		Head: &Node{
			Data: nil,
			Next: nil,
			Prev: nil,
		},
		errs: map[string]error{
			"OutOfBoundInsertion": fmt.Errorf("cannot insert at an out of bounds index into the list or list is empty. Use append() instead"),
			"OutOfBoundGeneric":   fmt.Errorf("out of bounds"),
		},
	}
}
