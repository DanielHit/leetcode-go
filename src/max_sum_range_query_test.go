package src

import (
	"fmt"
	"sort"
	"testing"
)

func Test_maxSumRangeQuery(t *testing.T) {
	countMap := map[int]int{}
	countMap[1] += 1
	countMap[1] += 1
	fmt.Println(countMap)
}

func Test_maxSumRangeQuery1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	requests := [][]int{{1, 3}, {0, 1}}
	maxSumRangeQuery(nums, requests)

	nums = []int{1, 2, 3, 4, 5, 6}
	requests = [][]int{{0, 1}}
	maxSumRangeQuery(nums, requests)

	nums = []int{1, 2, 3, 4, 5, 10}
	requests = [][]int{{0, 2}, {1, 3}, {1, 1}}
	maxSumRangeQuery(nums, requests)
}

func TestSort(t *testing.T) {
	nums := []int{23, 2, 45, 6, 89}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
}
