package main

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	data  []int
	size  int
	front int
	rear  int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, n),
		size:  n,
		front: -1,
		rear:  -1,
	}
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.front == -1
}

func (cq *CircularQueue) IsFull() bool {
	return (cq.rear+1)%cq.size == cq.front
}

func (cq *CircularQueue) Enqueue(x int) error {
	if cq.IsFull() {
		return errors.New("Queue is full")
	}

	cq.rear = (cq.rear + 1) % cq.size
	if cq.front == -1 {
		cq.front = 0
	}
	fmt.Println("cq.rear", cq.rear, "cq.front", cq.front)
	cq.data[cq.rear] = x
	return nil
}

func (cq *CircularQueue) Dequeue() (int, error) {
	if cq.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	x := cq.data[cq.front]
	if cq.front == cq.rear {
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % cq.size
	}

	return x, nil
}

func main() {
	queueSize := 5
	cq := NewCircularQueue(queueSize)
	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)
	cq.Enqueue(4)
	cq.Enqueue(5)

	// cq.Dequeue()
	cq.Dequeue()
	// // cq.Dequeue()
	// cq.Enqueue(10)
	// cq.Dequeue()
	// fmt.Println(cq.Enqueue(101))
	// cq.Dequeue()
	// fmt.Println(cq.Enqueue(6))
	// fmt.Println(cq.Enqueue(6))
	// cq.Dequeue()
	// fmt.Println(cq.Enqueue(7))

	// Enqueue elements
	// for i := 1; i <= 6; i++ {
	// 	if err := cq.Enqueue(i); err != nil {
	// 		cq.Dequeue()
	// 		fmt.Println(err)
	// 	}
	// }
	fmt.Printf("%+v \n", cq)
	// Dequeue and display elements
	// for i := 0; i < queueSize; i++ {
	// 	item, err := cq.Dequeue()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Printf("%d ", item)
	// 	}
	// }
	fmt.Println()
}
