package main

// BinarySearch implementation only returns bool if needle is found
func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1

	for low < high {
		mid := low + (high-low)/2
		guess := haystack[mid]

		// if correct
		if needle == guess {
			return true

			// if needle is higher
		} else if needle > guess {

			// + 1 is for dropping the midpoint
			low = mid + 1

			// if needle is lower
		} else if needle < guess {
			high = mid

		}
	}

	return false
}
