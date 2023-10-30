package sort

/*
	Bubble Sort: o(n^2) ---> Quadratic time complexity. (NESTED LOOOOOOP)
	This is a simple comparison-based sorting algorithm.
	It repeatedly steps through the list, compares adjacent elements,
	and swaps them if they are in the wrong order.
*/

func BubbleSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				swapped = true
			}
		}
	}
}
