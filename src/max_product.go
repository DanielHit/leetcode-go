package src

// 两个子串的回文数最大化
func maxProduct(s string) int64 {
	max := int64(0)
	for i := 0; i < len(s); i++ {
		left := s[:i]
		right := s[i:]
		temp := theMaxOddLengthPalindromic(left) * theMaxOddLengthPalindromic(right)
		if temp > max {
			max = temp
		}
	}
	return max
}

// "abacd" is 3
// marcher的算法 O(n)的时间复杂度.
func theMaxOddLengthPalindromic(s string) int64 {
	// 如何思考个人的最长子串.
	// thinking about the psychology of money
	if len(s) <= 1 {
		return int64(len(s))
	}

	// 判断是否是
	max := 1
	for middle := 1; middle < len(s); middle++ {
		step := 1
		for ; middle-step >= 0 && middle+step <= len(s)-1; step++ {
			if s[middle-step] == s[middle+step] {
				if step*2+1 > max {
					max = step*2 + 1
				}
			} else {
				break
			}
		}
	}

	return int64(max)
}
