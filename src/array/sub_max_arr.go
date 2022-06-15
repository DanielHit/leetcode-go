package array

import "math"

// calculate the sub arr max
func subMaxArray(arr []int) int {
	res := make([]int, len(arr))
	max := math.MinInt32
	res[0] = arr[0]
	sum := arr[0]
	for i := 1; i < len(arr); i++ {
		sum = maxF(sum+arr[i], arr[i])
		max = maxF(sum, max)
	}
	return max
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}
