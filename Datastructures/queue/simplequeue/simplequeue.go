package simpleQueu

import "fmt"

/*
	A Queue is a linear data structure and it follows FIFO.

	1.Enqueue: Add an element to the end of the queue
	2.Dequeue: Remove an element from the front of the queue
	3.IsEmpty: Check if the queue is empty
	4.IsFull: Check if the queue is full
	5.Peek: Get the value of the front of the queue without removing it
*/

type Queue struct {
	values   []int
	capacity int
}

func (q *Queue) enQueue(v int) {
	if len(q.values) >= q.capacity {
		panic("overflow!")
	}
	q.values = append(q.values, v)
}

func (q *Queue) deQueue() int {
	if q.isEmpty() {
		panic("empty queue!")
	}
	removedValue := q.values[0]
	q.values = q.values[1:]
	return removedValue
}

func (q *Queue) isEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue) peek() int {
	if q.isEmpty() {
		panic("empty queue!")
	}
	return q.values[0]
}

func SimpleQueu() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recovered: ", err)
		}
	}()
	q := &Queue{values: []int{}, capacity: 5}
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	fmt.Println(q)
	fmt.Println(q.deQueue())
	fmt.Println(q)
	fmt.Println(q.peek())
	q.enQueue(4)
	q.enQueue(5)
	q.enQueue(6)
	fmt.Println(q)
	q.enQueue(6)
}
