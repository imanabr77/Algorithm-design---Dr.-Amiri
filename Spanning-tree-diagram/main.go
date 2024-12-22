package main

import (
	"container/heap"
	"fmt"
)

// Edge represents a graph edge.
type Edge struct {
	to     int
	weight int
}

// PriorityQueueItem represents an item in the priority queue.
type PriorityQueueItem struct {
	node   int
	weight int
	index  int
}

// PriorityQueue implements a min-heap priority queue.
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].weight < pq[j].weight }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// PrimMST calculates the Minimum Spanning Tree using Prim's algorithm
func PrimMST(graph map[int][]Edge, start int) (int, []Edge) {
	n := len(graph)
	visited := make([]bool, n)
	mst := []Edge{}
	totalWeight := 0

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &PriorityQueueItem{node: start, weight: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*PriorityQueueItem)
		if visited[item.node] {
			continue
		}

		visited[item.node] = true
		totalWeight += item.weight
		if item.weight != 0 {
			mst = append(mst, Edge{to: item.node, weight: item.weight})
		}

		for _, edge := range graph[item.node] {
			if !visited[edge.to] {
				heap.Push(pq, &PriorityQueueItem{node: edge.to, weight: edge.weight})
			}
		}
	}

	return totalWeight, mst
}

func main() {
	// Define the graph as an adjacency list.
	graph := map[int][]Edge{
		0: {{to: 1, weight: 4}, {to: 6, weight: 7}},
		1: {{to: 0, weight: 4}, {to: 2, weight: 9}, {to: 6, weight: 11}, {to: 4, weight: 20}},
		2: {{to: 1, weight: 9}, {to: 3, weight: 6}, {to: 4, weight: 2}},
		3: {{to: 2, weight: 6}, {to: 4, weight: 10}, {to: 5, weight: 4}},
		4: {{to: 1, weight: 20}, {to: 2, weight: 2}, {to: 3, weight: 10}, {to: 5, weight: 15}, {to: 7, weight: 1}},
		5: {{to: 3, weight: 4}, {to: 4, weight: 15}, {to: 8, weight: 12}},
		6: {{to: 0, weight: 7}, {to: 1, weight: 11}, {to: 7, weight: 1}},
		7: {{to: 4, weight: 1}, {to: 6, weight: 1}, {to: 8, weight: 3}},
		8: {{to: 7, weight: 3}, {to: 5, weight: 12}},
	}

	totalWeight, mst := PrimMST(graph, 0)

	fmt.Printf("Total Weight of MST: %d\n", totalWeight)
	fmt.Println("Edges in the MST:")
	for _, edge := range mst {
		fmt.Printf("%d - %d (Weight: %d)\n", edge.to, edge.weight, edge.weight)
	}
}
