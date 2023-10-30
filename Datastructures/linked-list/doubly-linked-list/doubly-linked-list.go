package main

import "fmt"

// Node represents a node in the doubly linked list.
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// Append adds a new node with the given data to the end of the list.
func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}

	dll.Length++
}

// DisplayForward prints the elements of the linked list from head to tail.
func (dll *DoublyLinkedList) DisplayForward() {
	current := dll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

// DisplayBackward prints the elements of the linked list from tail to head.
func (dll *DoublyLinkedList) DisplayBackward() {
	current := dll.Tail
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Prev
	}
	fmt.Println("nil")
}

func main() {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	fmt.Println("Doubly Linked List (Forward):")
	dll.DisplayForward()

	fmt.Println("Doubly Linked List (Backward):")
	dll.DisplayBackward()
}
