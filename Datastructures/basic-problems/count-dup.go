package basicproblems

func findAccurances(arr [][]int, target int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == target {
				count++
			}
		}
	}
	return count
}
