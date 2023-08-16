package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	test := []int{1, 2, 4, 7, 6, 3, 5}
	got := BubbleSort(test)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("WRONG")
		t.Errorf("GOT %v", got)
		t.Errorf("WANT %v", want)
	}

	test = []int{7, 6, 5, 4, 3, 2, 1}
	got = BubbleSort(test)
	want = []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("WRONG")
		t.Errorf("GOT %v", got)
		t.Errorf("WANT %v", want)
	}
}
