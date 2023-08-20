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
	Prepend(item any)
	Remove(item any) error
	RemoveAt(index int) error
}

// LinkedList implementation
type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
	errs   CustomErrs
}

func (l *LinkedList) GetLength() int {
	return l.Length
}

// InsertAt inserts any data at any index within the list
func (l *LinkedList) InsertAt(item any, index int) error {
	node := l.Head

	if index > l.Length {
		return l.errs.OutOfBoundInsertion
	} else if l.Length == 0 {
		return l.errs.OutOfBoundInsertion
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
			node.Prev = current
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
	return nil, l.errs.OutOfBoundGeneric
}

func (l *LinkedList) Prepend(item any) {
	newHead := &Node{
		Data: item,
		Next: l.Head,
		Prev: nil,
	}
	l.Head.Prev = newHead
	l.Head = newHead
}

func (l *LinkedList) Remove(item any) error {
	current := l.Head
	for {

		if current.Data == item && current == l.Head { // if first item matches
			l.Head = current.Next
			current.Next.Prev = nil
			current.Next = nil
			l.Length--
			return nil

		} else if current.Data == item && current == l.Tail { // if last node matches
			l.Tail = current.Prev
			current.Prev.Next = nil
			current.Prev = nil
			l.Length--
			return nil

		} else if current.Data == item { // somewhere in the middle

			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev

			// don't think this is necessary
			current.Next = nil
			current.Prev = nil
			current.Data = nil
			l.Length--
			return nil

		}

		// reached the end
		if current == l.Tail {
			return l.errs.NotFoundGeneric
		}

		current = current.Next
	}
}

func (l *LinkedList) RemoveAt(index int) error {
	current := l.Head

	for i := 0; i < l.Length-1; i++ {
		if index == 0 { // if first item matches
			l.Head = current.Next
			current.Next.Prev = nil
			current.Next = nil
			l.Length--
			return nil

		} else if index == l.Length-1 { // if last node matches
			l.Tail = current.Prev
			current.Prev.Next = nil
			current.Prev = nil
			l.Length--
			return nil

		} else if index == i { // somewhere in the middle

			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev

			// don't think this is necessary
			current.Next = nil
			current.Prev = nil
			current.Data = nil
			l.Length--
			return nil
		}

		// reached the end
		if current == l.Tail {
			return l.errs.NotFoundGeneric
		}

		current = current.Next
	}

	return l.errs.NotFoundGeneric
}

type CustomErrs struct {
	OutOfBoundInsertion error
	OutOfBoundGeneric   error
	NotFoundGeneric     error
}

func Errs() CustomErrs {

	return CustomErrs{
		OutOfBoundGeneric:   fmt.Errorf("ERROR: out of bounds"),
		OutOfBoundInsertion: fmt.Errorf("ERROR: cannot insert at an out of bounds index into the list or list is empty. Use append() instead"),
		NotFoundGeneric:     fmt.Errorf("ERROR: item was not found"),
	}
}

func NewLinkedList() LinkedList {

	return LinkedList{
		Length: 0,
		Head: &Node{
			Data: nil,
			Next: nil,
			Prev: nil,
		},
		errs: Errs(),
	}
}
