package main

func BubbleSort(unsorted []int) []int {

	for i := 0; i < len(unsorted)-1; i++ {

		for j := 0; j < len(unsorted)-1-i; j++ {
			// if left is larger than right, swap
			if unsorted[j] > unsorted[j+1] {
				smaller := unsorted[j+1]
				larger := unsorted[j]
				unsorted[j] = smaller
				unsorted[j+1] = larger
			}
		}

	}

	// now sorted
	return unsorted
}

// QuickSort divide & conquer
func QuickSort() {

}
