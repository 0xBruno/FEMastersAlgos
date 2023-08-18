package main

func BubbleSort(unsorted []int) []int {

	// decrementing array
	for length := len(unsorted) - 1; length != 0; length-- {

		// loop over entire shrinking array
		for i := 0; i < length; i++ {

			// if left is larger than right, swap
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
