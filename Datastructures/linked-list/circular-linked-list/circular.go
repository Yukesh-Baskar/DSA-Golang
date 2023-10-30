package main

import "fmt"

// Node represents a node in the circular linked list.
type Node struct {
	Data int
	Next *Node
}

// CircularLinkedList represents a circular linked list.
type CircularLinkedList struct {
	Head *Node
	Tail *Node
}

// Append adds a new node with the given data to the end of the circular linked list.
func (cll *CircularLinkedList) Append(data int) {
	newNode := &Node{Data: data}

	if cll.Head == nil {
		cll.Head = newNode
		cll.Tail = newNode
		newNode.Next = newNode // Point to itself to create a circular reference
	} else {
		newNode.Next = cll.Head
		cll.Tail.Next = newNode
		cll.Tail = newNode
	}
}

// Display prints the elements of the circular linked list.
func (cll *CircularLinkedList) Display() {
	if cll.Head == nil {
		fmt.Println("Circular linked list is empty.")
		return
	}

	current := cll.Head
	for {
		fmt.Printf("%d __ %+v -> __ %+v -=->", current.Data, current.Next, cll.Tail)
		current = current.Next
		if current == cll.Head {
			break
		}
	}
	fmt.Println("Head (back to start)")
}

func main() {
	cll := CircularLinkedList{}

	cll.Append(1)
	cll.Append(2)
	cll.Append(3)

	fmt.Println("Circular Linked List:")
	cll.Display()
}
