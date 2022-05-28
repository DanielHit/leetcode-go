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
