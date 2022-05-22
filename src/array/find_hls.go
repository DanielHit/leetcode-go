package array

import (
	"sort"
)

// problem: 594. Longest Harmonious Subsequence
func findLHS(nums []int) int {
	sort.Ints(nums)

	tempMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		tempMap[nums[i]]++
	}

	before := -100000000
	beforeValue := 0
	count := 0

	for _, key := range nums {
		if key-before == 1 {
			if tempMap[key]+beforeValue > count {
				count = tempMap[key] + beforeValue
			}
		}
		before = key
		beforeValue = tempMap[key]
	}

	return count
}

func findLHS_V2(nums []int) int {

	tempMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		tempMap[nums[i]]++
	}

	count := 0
	for key, value := range tempMap {
		if v2, ok := tempMap[key+1]; ok {
			if value+v2 > count {
				count = value + v2
			}
		}
	}

	return count
}

func findLHS_V3(nums []int) int {

	tempMap := map[int]int{}
	count := 0
	for i := 0; i < len(nums); i++ {
		tempMap[nums[i]]++
		if v, ok := tempMap[nums[i]+1]; ok {
			if v+tempMap[nums[i]] > count {
				count = v + tempMap[nums[i]]
			}
		}

		if v, ok := tempMap[nums[i]-1]; ok {
			if v+tempMap[nums[i]] > count {
				count = v + tempMap[nums[i]]
			}
		}

	}

	return count
}
