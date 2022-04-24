package src

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {

	ori := make([]int, len(nums))
	copy(ori, nums)

	// 排序
	sort.Ints(nums)

	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i]+nums[j] == target {
			break
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}

	fmt.Println(nums, nums[i], nums[j])

	//
	first := -1
	second := -1
	for k := 0; k < len(ori); k++ {

		// 过渡
		if first == -1 && ori[k] == nums[i] {
			first = k
			continue
		}

		if second == -1 && ori[k] == nums[j] {
			second = k
			continue
		}

	}

	// 打印相应的知识点
	return []int{first, second}
}
