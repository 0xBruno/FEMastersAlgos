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
	EmptyQueue          error
	EmptyStack          error
}

func NewLinkedList() LinkedList {

	return LinkedList{
		Length: 0,
		Head: &Node{
			Data: nil,
			Next: nil,
			Prev: nil,
		},
		errs: CustomErrs{
			OutOfBoundGeneric:   fmt.Errorf("ERROR: out of bounds"),
			OutOfBoundInsertion: fmt.Errorf("ERROR: cannot insert at an out of bounds index into the list or list is empty. Use append() instead"),
			NotFoundGeneric:     fmt.Errorf("ERROR: item was not found"),
		},
	}
}

// QueueInterface api
type QueueInterface interface {
	Enqueue(item any)
	Dequeue() (any, error)
	Peek() any
}

type Queue struct {
	LinkedList
}

func NewQueue() Queue {
	return Queue{
		LinkedList{
			Length: 0,
			Head:   nil,
			Tail:   nil,
			errs: CustomErrs{
				EmptyQueue: fmt.Errorf("queue is empty! cannot dequeue"),
			},
		},
	}
}

func (l *Queue) Enqueue(item any) {
	l.Append(item)
}

func (l *Queue) Dequeue() (any, error) {
	if l.Head == nil {
		return nil, l.errs.EmptyQueue
	}

	// if queue only has one node
	if l.Head.Next == nil {
		l.Length--
		return l.Head.Data, nil
	}

	oldHead := l.Head
	l.Head = oldHead.Next
	oldHead.Next.Prev = nil
	oldHead.Next = nil
	l.Length--
	return oldHead.Data, nil
}

func (l *Queue) Peek() any {
	return l.Head.Data
}

// StackInterface api
type StackInterface interface {
	Push(item any)
	Pop() (any, error)
	Peek() any
}

type Stack struct {
	LinkedList
}

func NewStack() Stack {
	return Stack{
		LinkedList{
			Length: 0,
			Head:   nil,
			Tail:   nil,
			errs: CustomErrs{
				EmptyStack: fmt.Errorf("stack is empty! cannot pop element"),
			},
		},
	}
}

func (l *Stack) Push(item any) {
	l.Append(item)
}

func (l *Stack) Pop() (any, error) {

	if l.Tail == nil {
		return nil, l.errs.EmptyStack
	}

	popped := l.Tail
	prev := l.Tail.Prev
	// remove forward link
	prev.Next = nil
	// remove back link
	popped.Prev = nil
	// update ll
	l.Tail = prev
	l.Length--
	return popped.Data, nil
}

func (l *Stack) Peek() any {
	return l.Tail.Data
}

type MinHeapInterface interface {
	Insert(value int)
	Remove() (int, error)
}

type MinHeap struct {
	data   []int
	Length int
}

func (m *MinHeap) Insert(value int) {
	m.data = append(m.data, value)
	m.heapifyUp(m.Length)
	m.Length++

}

func (m *MinHeap) Remove() (int, error) {
	if m.Length == 0 {
		return -1, fmt.Errorf("heap is empty")
	}

	out := m.data[0]
	m.Length--

	if m.Length == 0 {
		m.data = []int{}
		return out, nil
	}

	m.data[0] = m.data[m.Length]
	m.heapifyDown(0)

	return out, nil
}

func (m *MinHeap) heapifyDown(index int) {

	lIdx := m.leftChild(index)
	rIdx := m.rightChild(index)

	// base case
	// since we are always working L to R we only check Left
	if index >= m.Length || lIdx >= m.Length {
		return
	}

	lVal := m.data[lIdx]
	rVal := m.data[rIdx]
	currVal := m.data[index]

	if lVal > rVal && currVal > rVal {
		m.data[index] = rVal
		m.data[rIdx] = currVal
		m.heapifyDown(rIdx)
	} else if rVal > lVal && currVal > lVal {
		m.data[index] = lVal
		m.data[lIdx] = currVal
		m.heapifyDown(lIdx)
	}

}

func (m *MinHeap) heapifyUp(index int) {

	if index == 0 {
		return
	}

	parent := m.parent(index)
	parentValue := m.data[parent]
	currValue := m.data[index]

	if parentValue > currValue {
		// swap current with parent
		m.data[index] = parentValue
		m.data[parent] = currValue
		m.heapifyUp(parent)
	}

}

func (m *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *MinHeap) leftChild(index int) int {
	return 2*index + 1
}

func (m *MinHeap) rightChild(index int) int {
	return 2*index + 2
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data:   []int{},
		Length: 0,
	}
}

type LRUNode struct {
	value any
	next  *LRUNode
	prev  *LRUNode
}

type LRUInterface interface {
	Update(key string, value any)
	Get(key string) error
}

type LRU struct {
	length   int
	head     *LRUNode
	tail     *LRUNode
	capacity int

	lookup        map[string]*LRUNode
	reverseLookup map[*LRUNode]string
}

func (l *LRU) Update(key string, value any) {
	node, err := l.get(key)

	if err != nil {
		node = NewLRUNode(value)
		l.length++
		l.prepend(node)
		l.trimCache()

		l.lookup[key] = node
		l.reverseLookup[node] = key

	} else {
		l.detach(node)
		l.prepend(node)
		node.value = value
	}
}

// internal implementation that returns an *LRUNode
func (l *LRU) get(key string) (*LRUNode, error) {
	// check the cache for existence
	node := l.lookup[key]

	if node == nil {
		return nil, fmt.Errorf("item not found")
	}

	// update the value we found and move it to the front
	l.detach(node)
	l.prepend(node)

	return node, nil
}

// Get external implementation that returns the LRUNode.value
func (l *LRU) Get(key string) (any, error) {
	// check the cache for existence
	node := l.lookup[key]

	if node == nil {
		return nil, fmt.Errorf("item not found")
	}

	// update the value we found and move it to the front
	l.detach(node)
	l.prepend(node)

	return node.value, nil
}

func (l *LRU) detach(node *LRUNode) {
	// take node from behind and point it to node ahead
	if node.prev != nil {
		node.prev.next = node.next
	}

	// take node from ahead and point it to behind
	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.length == 1 {
		l.tail = nil
		l.head = nil
	}

	if l.head == node {
		l.head = l.head.next
	}

	if l.tail == node {
		l.tail = l.tail.prev
	}
	// remove links from current node
	node.prev = nil
	node.next = nil
}

func (l *LRU) prepend(node *LRUNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *LRU) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)
	key := l.reverseLookup[tail]

	delete(l.lookup, key)
	delete(l.reverseLookup, l.tail)
	l.length--

}

func NewLRU(capacity int) *LRU {
	return &LRU{
		length:        0,
		head:          nil,
		tail:          nil,
		capacity:      capacity,
		lookup:        map[string]*LRUNode{},
		reverseLookup: map[*LRUNode]string{},
	}
}

func NewLRUNode(val any) *LRUNode {
	return &LRUNode{
		value: val,
	}
}
