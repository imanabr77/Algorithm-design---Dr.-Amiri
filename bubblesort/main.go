package main

import "fmt"

// BubbleSort function to sort an array
func BubbleSort(arr []int) {
	n := len(arr)

	// Traverse through all elements in the array
	for i := 0; i < n-1; i++ {
		// Last i elements are already sorted, so skip them
		for j := 0; j < n-i-1; j++ {
			// Swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Example array to sort
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted array:", arr)

	// Sort the array
	BubbleSort(arr)

	fmt.Println("Sorted array:", arr)
}
