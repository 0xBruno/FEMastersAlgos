package main

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	// putting it all in one func to avoid fixtures
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

	// get()
	list.Prepend("FIRST!")
	got = list.Head.Data
	want = "FIRST!"

	if got != want {
		t.Errorf(got.(string), want)
	}

	err := list.Remove("FIRST!")

	if err != nil {
		t.Fatal(err)
	}

	got = list.Head.Data
	want = "LL 0"

	if got != want {
		t.Errorf(got.(string), want)
	}

	err = list.Remove("LL 3")

	if err != nil {
		t.Fatal(err)
	}

	// would get LL 3 if not removed
	got = list.Head.Next.Next.Next.Data
	want = "LL 4"

	if got != want {
		t.Errorf(got.(string), want)
	}

	err = list.Remove("LL 1000")

	if err != nil {
		t.Fatal(err)
	}

	got = list.Tail.Data
	want = "LL 999"

	if got != want {
		t.Errorf(got.(string), want)
	}

	// zeroth index is removed twice
	err = list.RemoveAt(0)
	err = list.RemoveAt(0)

	if err != nil {
		t.Fatal(err)
	}

	got = list.Head.Data
	want = "LL 2"

	if got != want {
		t.Errorf(got.(string), want)
	}

	fmt.Println(list.Length)
}
