package array

// jump solve problem: 45. Jump Game II
func jump(nums []int) int {
	// find the right solution to the answer

	count := 0
	start := len(nums) - 1

	for {
		if start == 0 {
			break
		}

		temp := 0
		for i := start; i >= 0; i-- {
			if start-i <= nums[i] {
				temp = i
			}
		}
		start = temp
		count++
	}

	return count
}
