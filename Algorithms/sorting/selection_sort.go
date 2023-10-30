package sort

import (
	"fmt"
	"time"
)

/*
	Selection Sort:  o(n^2) ---> Quadratic time complexity. (NESTED LOOOOOOP)
	Selection sort is another simple sorting algorithm.
	It divides the input list into two parts: the left subarray, which is sorted,
	and the right subarray, which is unsorted.
	It repeatedly selects the minimum element
	from the unsorted portion and moves it to the sorted portion.
	[4,5,2,1,6,3,0,7]
*/

func SelectionSort(nums []int) {
	start := time.Now()
	defer func() {
		fmt.Println("time taken: ", time.Since(start))
	}()
	l := len(nums)
	if l == 0 {
		return
	}
	for i := 0; i < l-1; i++ {
		minIndex := i

		for j := i + 1; j < l; j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			nums[minIndex], nums[i] = nums[i], nums[minIndex]
		}
	}
}
