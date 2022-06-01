package array

// problem: 1480. Running Sum of 1d Array
func runningSum(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}
