package main

import (
	"container/heap"
	"fmt"
)

// Item represents an item with a priority.
type Item struct {
	value    string
	priority int
}

// PriorityQueue implements a priority queue.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	items := []*Item{
		{value: "Item 1", priority: 3},
		{value: "Item 2", priority: 1},
		{value: "Item 3", priority: 4},
		{value: "Item 4", priority: 2},
	}

	// Push items onto the priority queue
	for _, item := range items {
		heap.Push(&pq, item)
	}

	// Pop items from the priority queue in order of priority
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.value, item.priority)
	}
}
