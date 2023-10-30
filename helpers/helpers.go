package main

import "fmt"

func main() {
	// *[]int: A pointer to a slice of integers
	slice1 := []int{1, 2, 3}
	slicePtr1 := &slice1
	fmt.Println("Original slice1:", slice1)
	(*slicePtr1)[0] = 99
	fmt.Println("Modified slice1:", slice1)

	// *[]*int: A pointer to a slice of integer pointers
	value1 := 42
	value2 := 54
	value3 := 67
	slice2 := []*int{&value1, &value2, &value3}
	slicePtr2 := &slice2
	fmt.Println("Original slice2:", *slice2[0], *slice2[1], *slice2[2])
	*(*slicePtr2)[0] = 100
	fmt.Println("Modified slice2:", *slice2[0], *slice2[1], *slice2[2])
}
