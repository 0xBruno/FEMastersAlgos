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

func partition(arr *[]int, low int, high int) int {
	pivot := (*arr)[high]
	index := low - 1

	for i := low; i < high; i++ {
		if (*arr)[i] <= pivot {
			index++
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[index]
			(*arr)[index] = tmp
		}
	}

	// swap pivot with pivot index
	index++
	(*arr)[high] = (*arr)[index]
	(*arr)[index] = pivot

	return index
}

func qs(arr *[]int, low int, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)

	qs(arr, low, pivotIndex-1)
	qs(arr, pivotIndex+1, high)

}

// QuickSort divide & conquer
func QuickSort(arr *[]int) {
	// this one was definitely a straight port over and need to study it more
	// two steps
	// 1. partition
	// 2. quick sort
	qs(arr, 0, len(*arr)-1)
}
