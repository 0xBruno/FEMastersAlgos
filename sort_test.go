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

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	QuickSort(&arr)

	if !reflect.DeepEqual(arr, []int{3, 4, 7, 9, 42, 69, 420}) {
		t.Fatal(arr)
	}

}
