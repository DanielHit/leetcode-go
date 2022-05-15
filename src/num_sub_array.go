package src

// https://leetcode.com/problems/binary-subarrays-with-sum/
func numSubarraysWithSum(nums []int, goal int) int {
	res := 0
	var path []int
	recurseSubArray(nums, path, &res, 0, goal)
	return res
}

func recurseSubArray(nums []int, path []int, res *int, start int, goal int) {
	if sum(path) == goal {
		*res += 1
		return
	}
	if start == len(nums)-1 && sum(path) != goal {
		return
	}

	if start >= len(nums) || len(path) > len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		recurseSubArray(nums, path, res, start+1, goal)
		path = path[:len(path)-1]
	}
}

func sum(path []int) int {
	s := 0
	for _, v := range path {
		s += v
	}
	return s
}
