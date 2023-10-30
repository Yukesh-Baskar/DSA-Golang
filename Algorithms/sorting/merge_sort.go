package sort

import "fmt"

/*
	Merge Sort: O(n log n)
	Merge Sort is a divide-and-conquer algorithm
	that works by dividing the unsorted list into several
	sublists until each sublist contains only one element.
	It then merges these sublists in a manner that combines them in a sorted fashion.

*/

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		// fmt.Println("arrR", arr)
		return arr
	}

	// Divide the array into two halves
	mid := len(arr) / 2
	fmt.Println("arr", arr, "leftHalf: ", arr[:mid], "rightHalf: ", arr[mid:], mid)
	leftHalf := MergeSort(arr[:mid])
	fmt.Println("leftHalf", leftHalf)
	rightHalf := MergeSort(arr[mid:])
	fmt.Println("rightHalf", rightHalf)
	// Merge the sorted halves
	return Merge(leftHalf, rightHalf)
}

func Merge(leftHalf, rightHalf []int) []int {
	fmt.Println("leftHalf, rightHalf", leftHalf, rightHalf)
	result := make([]int, 0, len(leftHalf)+len(rightHalf))
	i, j := 0, 0

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] <= rightHalf[j] {
			result = append(result, leftHalf[i])
			i++
		} else {
			result = append(result, rightHalf[j])
			j++
		}
	}

	result = append(result, leftHalf[i:]...)
	result = append(result, rightHalf[j:]...)
	fmt.Println("result", result, i, j, leftHalf[i:], rightHalf[j:])
	return result
}

// [15,21] [2,5,4] SHOULD PASS ALREADY SHORTED ARRAY
func SortTwoArrays(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	fmt.Println("res", result) // Debugging print statement
	return result
}
