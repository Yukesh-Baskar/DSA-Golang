package search

// Loops through every element in an array. 
// LINEAR TIME COMPLEXITY ----> o(n)

func LinearSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}
