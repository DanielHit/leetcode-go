package src

// loop for sum
// test the middle class
// 2022-04-25
func findMiddleIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftPartSum := 0
	index := 0
	for ; index < len(nums); index++ {
		if sum-nums[index] == leftPartSum*2 {
			return index
		}
		leftPartSum += nums[index]
	}

	return -1
}
