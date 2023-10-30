package search

// [1,2,3,4,5,6,7,8,9,10] -> 4
func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1

	for i := 0; i <= len(arr)/2; i++ {
		mid := (left + right/2)

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
