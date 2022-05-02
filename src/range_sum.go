package src

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	MODULO := 1000000007
	resNum := []int{}
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := i; j < len(nums); j++ {
			temp += nums[j]
			resNum = append(resNum, temp)
		}
	}

	// 有溢出的风险
	sort.Ints(resNum)
	res := 0

	for i := left - 1; i <= right-1; i++ {
		res += resNum[i]
	}
	res = res % MODULO
	return res
}
