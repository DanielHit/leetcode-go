package dp

import "strconv"

func majority(arr []int) (map[string]int, map[string]int) {
	dp := map[string]int{}
	dpChar := map[string]int{}

	for i := 0; i < len(arr); i++ {
		countMap := map[int]int{}
		for j := i; j < len(arr); j++ {
			countMap[arr[j]]++
			if i == j {
				dp[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = 1
				dpChar[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = arr[i]
			} else {
				if arr[j] == dpChar[strconv.Itoa(i)+"_"+strconv.Itoa(j-1)] {
					dp[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = dp[strconv.Itoa(i)+"_"+strconv.Itoa(j-1)] + 1
					dpChar[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = arr[j]
				} else {
					count := countMap[arr[j]]
					if count > dp[strconv.Itoa(i)+"_"+strconv.Itoa(j-1)] {
						dp[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = count
						dpChar[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = arr[j]
					} else {
						dp[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = dp[strconv.Itoa(i)+"_"+strconv.Itoa(j-1)]
						dpChar[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = dpChar[strconv.Itoa(i)+"_"+strconv.Itoa(j-1)]
					}
				}
			}
		}
	}
	return dp, dpChar
}

func majorityII(arr []int) ([][]int, [][]int) {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}

	dpMap := make([][]int, len(arr))
	for i := range dp {
		dpMap[i] = make([]int, len(arr))
	}

	for i := 0; i < len(arr); i++ {
		countMap := map[int]int{}
		for j := i; j < len(arr); j++ {
			if j == i {
				dp[i][j] = 1
				dpMap[i][j] = arr[j]
			} else if arr[j] == dpMap[i][j-1] {
				dp[i][j] = dp[i][j-1] + 1
				dpMap[i][j] = arr[j]
			} else {
				if countMap[arr[j]]+1 > dpMap[i][j-1] {
					dp[i][j] = countMap[arr[j]] + 1
					dpMap[i][j] = arr[j]
				} else {
					dpMap[i][j] = dpMap[i][j-1]
					dp[i][j] = dp[i][j-1]
				}
			}

			countMap[arr[j]]++
		}
	}
	return dp, dpMap
}
