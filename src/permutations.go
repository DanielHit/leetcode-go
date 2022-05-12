package src

import "fmt"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	recursive(&res, 0, &nums)
	fmt.Println("hell:", nums)
	return res
}

func recursive(res *[][]int, start int, nums *[]int) {
	if start >= len(*nums) {
		*res = append(*res, *nums)
		return
	}

	for i := start; i < len(*nums); i++ {
		swap(*nums, start, i)
		recursive(res, start+1, nums)
		swap(*nums, i, start)
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
