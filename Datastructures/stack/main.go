package main

import "fmt"

/*
	A stack is a non linear data structure which will be stored in non-contigious memory location.
	Stack follows the LIFO(Last In First Out) principle (eg : plates ordered horizontally).

	These are all the methods a stack has:
	1.PUSH
	2.POP
	3.ISFULL
	4.PEEL
	5.ISEMPTY
*/

type Stack struct {
	values   []int
	capacity int
}

func (s *Stack) push(v int) {
	if s.isFull() {
		panic("stack overflow!")
	}
	s.values = append(s.values, v)
}

func (s *Stack) pop() {
	if s.isEmpty() {
		panic("pop on empty stack")
	}
	s.values = s.values[:len(s.values)-1]
}

func (s *Stack) peek() int {
	if s.isEmpty() {
		panic("empty stack, no peek values")
	}
	return s.values[0]
}
func (s *Stack) isFull() bool {
	if len(s.values) >= s.capacity {
		return true
	}
	return false
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func main() {
	s := &Stack{values: []int{}, capacity: 5}
	fmt.Println("before", s.isEmpty())
	s.push(1)
	fmt.Println("after", s.isEmpty())
	s.push(2)
	s.push(3)
	s.push(4)

	s.push(5)
	fmt.Println(s)
	fmt.Println(s.peek())
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()

	// s.push(6)
	fmt.Println(s)
}
