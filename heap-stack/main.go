package main

import (
	"container/heap"
	"fmt"
)

// Example of stack implementation in Go
type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Example of heap implementation in Go using a min-heap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func main() {
	// Stack example
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Stack:")
	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println(value)
	}

	// Heap example
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, 10)
	heap.Push(h, 5)
	heap.Push(h, 20)
	heap.Push(h, 15)
	fmt.Println("Heap:")
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
