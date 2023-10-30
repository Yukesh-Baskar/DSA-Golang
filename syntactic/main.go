package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type isHomeless bool

type Address struct {
	DoorNumber int
	City       string
	Country    string
	PostalCode int
}

type CustomMap[T comparable, V constraints.Ordered | ~bool] map[T]V

type CustomData interface {
	constraints.Ordered | ~bool | Address
}

type User[T CustomData] struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Data      T
}

func (u *User[CustomData]) getData() {
	fmt.Println(u.Data)
}

func recursive(n int) int {
	if n <= 1 {
		return n
	}
	fmt.Println("hello", n)
	num := recursive(n - 1)
	fmt.Println("num", num)
	return num * 2
}

// r(5)
// 	 r(4)
// 		r(3)
// 		 	r(2)
// 				r(1)
// 					return 1
// 			return 2
// 		return 4
// 	 return 8
// return 16

func main() {
	fmt.Println("res", recursive(5))
	siva := &User[isHomeless]{
		FirstName: "Sivasankaran",
		LastName:  "WaterMenon",
		Age:       99,
		Gender:    "Trans",
		Data:      true,
	}

	arun := &User[Address]{
		FirstName: "Arun",
		LastName:  "Kumar",
		Age:       99,
		Gender:    "Trans",
		Data: Address{
			DoorNumber: 21,
			City:       "Bengaluru",
			Country:    "India",
			PostalCode: 124324,
		},
	}
	siva.getData()
	arun.getData()

	m := make(CustomMap[int, string])
	m[1] = "2"
	m2 := make(CustomMap[float32, bool])
	m2[2.2] = true
	fmt.Println(m)
	fmt.Println(m2)
}
