package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	currNode := l.Head

	for currNode.Next != nil {
		currNode = currNode.Next
	}

	currNode.Next = newNode
}

func (l *LinkedList) Display() {
	if l.Head == nil {
		fmt.Println("head is empty")
		return
	}

	currNode := l.Head

	for currNode != nil {
		fmt.Println(currNode.Data, currNode.Next)
		currNode = currNode.Next
	}
}

func (l *LinkedList) DeleteNode(data int) {
	if l.Head == nil {
		fmt.Println("Head is empty!")
		return
	}

	// Handle the case where the node to be deleted is the Head node.
	if l.Head.Data == data {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	currNode := l.Head
	prevNode := l.Head

	for currNode != nil {
		if currNode.Data == data {
			prevNode.Next = currNode.Next
			l.Length--
			fmt.Println("Node with Data", data, "deleted.")
			return
		}
		prevNode = currNode
		currNode = currNode.Next
	}

	fmt.Println("No node found with Data", data)
}

func main() {
	l := LinkedList{}
	l.append(1)
	l.append(2)
	l.append(3)
	l.append(4)
	l.Display()
	fmt.Println("befor")
	// l.DeleteNode(1)
	l.DeleteNode(1)
	fmt.Println("after")
	l.Display()
}
