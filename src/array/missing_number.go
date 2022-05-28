package array

// missingNumber: 268. Missing Number
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
