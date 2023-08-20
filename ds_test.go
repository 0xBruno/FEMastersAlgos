package main

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()

	for i := 0; i <= 1000; i++ {
		list.append(fmt.Sprintf("LL %d", i))
	}

	// test head
	got := list.head.data
	want := "LL 0"

	if got != want {
		t.Errorf(got.(string), want)
	}

	// test tail
	got = list.tail.data
	want = "LL 1000"

	if got != want {
		t.Errorf(got.(string), want)
	}

}
