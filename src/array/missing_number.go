package array

// missingNumber: 268. Missing Number
// Space Complexity: O(1) Runtime Complexity: O(n)
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
