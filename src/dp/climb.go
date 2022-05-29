package dp

import "math"

// 爬楼梯. 每次只爬 1个楼梯或者2个楼梯.
// 使用递归
var resMap map[int]int

func climb(nums []int) int {
	resMap = map[int]int{}
	return int(math.Min(float64(climbHelper(nums, len(nums)-1)), float64(climbHelper(nums, len(nums)-2))))
}

// use dynamic
func climbHelper(nums []int, index int) int {
	if index < 2 {
		return nums[index]
	}
	if v, ok := resMap[index]; ok {
		return v
	}
	v := min(climbHelper(nums, index-1), climbHelper(nums, index-2)) + nums[index]
	resMap[index] = v
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
