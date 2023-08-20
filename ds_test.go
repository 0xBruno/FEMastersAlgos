package main

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()

	for i := 0; i <= 1000; i++ {
		list.Append(fmt.Sprintf("LL %d", i))
	}

	// test head
	got := list.Head.Data
	want := "LL 0"

	if got != want {
		t.Errorf(got.(string), want)
	}

	// test tail
	got = list.Tail.Data
	want = "LL 1000"

	if got != want {
		t.Errorf(got.(string), want)
	}

	// get()
	got, _ = list.Get(69)
	want = "LL 69"

	if got != want {
		t.Errorf(got.(string), want)
	}
}
