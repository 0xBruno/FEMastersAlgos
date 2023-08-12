package main

import "testing"

func TestBinarySearch(t *testing.T) {
	got := BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)
	want := true

	if got != want {
		t.Errorf("Not found")
	}

	got = BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 42)
	want = false

	if got != want {
		t.Errorf("Not found")
	}

}
