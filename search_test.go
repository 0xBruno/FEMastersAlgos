package main

import "testing"

func TestBinarySearch(t *testing.T) {
	got := BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)
	if got != true {
		t.Errorf("Not found")
	}

	got = BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 42)

	if got != false {
		t.Errorf("Not found")
	}

}

func TestTwoCrystalBalls(t *testing.T) {
	got := TwoCrystalBalls([]bool{false, false, false, true, true})
	want := 3

	if got != want {

		t.Errorf("INDEX INCORRECT")
		t.Errorf("GOT %d", got)
		t.Errorf("WANT %d", want)
	}

	got = TwoCrystalBalls([]bool{true, true, true, true, true})
	want = 1

	if got != want {
		t.Errorf("INDEX INCORRECT")
		t.Errorf("GOT %d", got)
		t.Errorf("WANT %d", want)
	}

}
