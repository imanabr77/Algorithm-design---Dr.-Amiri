package main

import (
	"fmt"
)

// QuickSort function sorts an array using the QuickSort algorithm
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Choose a pivot (here, the middle element)
	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	middle := []int{}

	// Partition the array into three parts
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}

	// Recursively sort left and right parts, and merge
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, middle...), right...)
}

func main() {
	// Example array
	arr := []int{10, 7, 8, 1, 5, 9, 3, 2, 4, 6}
	fmt.Println("Unsorted array:", arr)

	// Perform QuickSort
	sortedArr := QuickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
