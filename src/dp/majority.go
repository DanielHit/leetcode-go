package dp

import "fmt"

//
func majority(arr []int) ([][]int, [][]int) {
	dp := make([][]int, len(arr))
	for index := range dp {
		dp[index] = make([]int, len(arr))
	}

	dpChar := make([][]int, len(arr))
	for index := range dp {
		dpChar[index] = make([]int, len(arr))
	}

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if i == j {
				dp[i][j] = 1
				dpChar[i][j] = arr[i]
			} else {
				if arr[j] == dpChar[i][j-1] {
					dp[i][j] = dp[i][j-1] + 1
					dpChar[i][j] = arr[j]
				} else {
					count := 0
					for k := i; k <= j; k++ {
						if arr[k] == arr[j] {
							count++
						}
					}
					if count > dp[i][j-1] {
						dp[i][j] = count
						dpChar[i][j] = arr[j]
					} else {
						dp[i][j] = dp[i][j-1]
						dpChar[i][j] = dpChar[i][j-1]
					}
				}
			}
		}
	}
	fmt.Println(dp)
	return dp, dpChar
}
