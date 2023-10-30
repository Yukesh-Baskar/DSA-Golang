package sort

// [5,2,4,3,1]
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		if arr[j+1] != key {
			arr[j+1] = key
		}
	}
}
