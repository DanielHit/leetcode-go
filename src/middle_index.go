package src

import "fmt"

// loop for sum
// test the middle class
// 2022-04-25
func findMiddleIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	fmt.Println(sum)

	temp := 0
	index := 0
	for ; index < len(nums); index++ {
		if sum-nums[index] == temp*2 {
			return index
		}
		temp += nums[index]
	}

	return -1
}
