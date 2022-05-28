package array

import (
	"math"
)

// problem 120. Triangle
func minimumTotal(triangle [][]int) int {
	var min int
	min = math.MaxInt64
	minimumTotalHelper(triangle, 0, []int{}, 0, &min)
	return min
}

// use dp method
func minimumTotalDp(triangle [][]int) int {
	var dp [][]int
	for i := 0; i < len(triangle); i++ {
		dp = append(dp, make([]int, len(triangle[i])))
	}

	// dp[index][i] = dp
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < len(triangle[i]); j++ {
			if j >= len(dp[i-1]) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				if dp[i-1][j-1] < dp[i-1][j] {
					dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				} else {
					dp[i][j] = dp[i-1][j] + triangle[i][j]
				}
			}
		}
	}

	min := math.MaxInt32
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		if dp[len(triangle)-1][i] < min {
			min = dp[len(triangle)-1][i]
		}
	}
	return min
}

func minimumTotalHelper(triangle [][]int, start int, path []int, sum int, min *int) {
	if len(path) == len(triangle) {
		if sum < *min {
			*min = sum
		}
		return
	}

	if len(path) == 0 {
		path = append(path, 0)
		sum += triangle[0][0]
		start += 1
		minimumTotalHelper(triangle, start, path, sum, min)
		return
	}

	i := path[len(path)-1]
	path = append(path, i)
	sum += triangle[start][i]
	minimumTotalHelper(triangle, start+1, path, sum, min)
	path = path[:len(path)-1]
	sum -= triangle[start][i]

	j := i + 1
	if j < len(triangle[start]) {
		path = append(path, j)
		sum += triangle[start][j]
		minimumTotalHelper(triangle, start+1, path, sum, min)
		path = path[:len(path)-1]
		sum -= triangle[start][j]
	}
}
