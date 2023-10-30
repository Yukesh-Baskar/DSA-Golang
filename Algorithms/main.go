package main

import (
	"fmt"
	"sorting/search"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// left := []int{12, 14, 21, 12, 2}
	// right := []int{1, 23, 22}
	// sort.BubbleSort(nums)
	// sort.SelectionSort(nums)
	// sort.InsertionSort(nums)
	// fmt.Println("nus", nums)
	// res := sort.SortTwoArrays([]int{15, 21, 51}, []int{2, 3, 5, 12})
	// fmt.Println(res)
	fmt.Println(search.BinarySearch(nums, 5))
}
