package main

func BubbleSort(unsorted []int) []int {

	for length := len(unsorted) - 1; length != 0; length-- {
		for i := 0; i < length; i++ {
			if unsorted[i] > unsorted[i+1] {
				smaller := unsorted[i+1]
				larger := unsorted[i]
				unsorted[i] = smaller
				unsorted[i+1] = larger
			}
		}
	}

	// now sorted
	return unsorted
}
