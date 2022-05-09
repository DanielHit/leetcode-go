package src

import (
	"sort"
)

// problem: https://leetcode.com/problems/maximum-sum-obtained-of-any-permutation/
func maxSumRangeQuery(nums []int, requests [][]int) int {
	sort.Ints(nums)
	countMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(requests); j++ {
			if i >= requests[j][0] && i <= requests[j][1] {
				countMap[i] += 1
			}
		}
	}

	var keys []int
	for i := range countMap {
		keys = append(keys, i)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return countMap[keys[i]] > countMap[keys[j]]
	})

	sum := 0
	for index, v := range keys {
		sum += countMap[v] * nums[len(nums)-1-index]
	}

	return sum % (1000000000 + 7)
}
