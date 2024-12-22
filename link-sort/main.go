package main

import (
	"fmt"
)

// Define a Node structure for the linked list
type Node struct {
	data int
	next *Node
}

// MergeSort is the main function to sort the linked list
func MergeSort(head *Node) *Node {
	// Base case: if the list is empty or has one element
	if head == nil || head.next == nil {
		return head
	}

	// Split the linked list into two halves
	mid := getMiddle(head)
	left := head
	right := mid.next
	mid.next = nil

	// Recursively sort both halves
	left = MergeSort(left)
	right = MergeSort(right)

	// Merge the sorted halves and return the merged list
	return merge(left, right)
}

// Helper function to merge two sorted linked lists
func merge(left, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	// Compare and merge the two lists
	if left.data < right.data {
		left.next = merge(left.next, right)
		return left
	}
	right.next = merge(left, right.next)
	return right
}

// Helper function to find the middle node of a linked list
func getMiddle(head *Node) *Node {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// Utility function to append a new node at the end of the list
func appendNode(head *Node, data int) *Node {
	newNode := &Node{data: data}
	if head == nil {
		return newNode
	}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return head
}

// Utility function to print the linked list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {
	// Creating an unsorted linked list
	var head *Node
	data := []int{64, 34, 25, 12, 22, 11, 90}

	// Append data to the linked list
	for _, val := range data {
		head = appendNode(head, val)
	}

	fmt.Println("Unsorted linked list:")
	printList(head)

	// Sort the linked list using Merge Sort
	head = MergeSort(head)

	fmt.Println("Sorted linked list:")
	printList(head)
}
