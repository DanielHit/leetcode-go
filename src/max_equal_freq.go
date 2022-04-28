package src

// https://leetcode.com/problems/maximum-equal-frequency/
func maxEqualFreq(nums []int) int {
	if isSubArrayEqualFreq(nums) {
		return len(nums)
	}

	for i := len(nums) - 1; i >= 1; i-- {
		// 变换子数组
		dstNums := nums[0:i]
		if isSubArrayEqualFreq(dstNums) {
			return i
		}
	}

	return 0
}

// 这个子数组是否是符合的数组(移除1个后属于完全的相等个数的序列)
func isSubArrayEqualFreq(nums []int) bool {
	countMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := countMap[nums[i]]; ok {
			countMap[nums[i]] += 1
		} else {
			countMap[nums[i]] = 1
		}
	}

	lenMap := map[int]int{}
	for _, v := range countMap {
		if _, ok := lenMap[v]; ok {
			lenMap[v] += 1
		} else {
			lenMap[v] = 1
		}
	}

	if len(lenMap) == 2 {
		return true
	}

	return false
}
