package rank

import (
	"fmt"
	"math"
)

// mergeSort: 归并排序 分治算法
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums) / 2
	left := mergeSort(nums[:middle])
	right := mergeSort(nums[middle:])

	i := 0
	j := 0

	var res []int
	for {
		if i > len(left)-1 || j > len(right)-1 {
			break
		}
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	for ; i <= len(left)-1; i++ {
		res = append(res, left[i])
	}
	for ; j <= len(right)-1; j++ {
		res = append(res, right[j])
	}

	return res
}

// selectSort: 插入排序
func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				swap(nums, i, j)
			}
		}
	}
}

// todo 统计排序. 统计比当前数字少的
func countSort(nums []int) {
	if len(nums) <= 0 {
		return
	}

	min := math.MaxInt32
	max := math.MinInt32
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(min, max)
	count := make([]int, max-min+1)
	fmt.Println(len(count))
	for _, v := range nums {
		count[v-min]++
	}

	i := 0
	for num := min; num <= max; num++ {
		for count[num-min] > 0 {
			nums[i] = num
			i++
			count[num-min]--
		}
	}

}

// bubbleSort: 冒泡排序
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
