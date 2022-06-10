package string

// 3. Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := 0
	for i := 0; i < len(s); i++ {
		cMap := map[string]bool{}
		for j := i; j < len(s); j++ {
			if cMap[string(s[j])] {
				break
			} else {
				if j-i+1 > max {
					max = j - i + 1
				}
				cMap[string(s[j])] = true
			}
		}
	}
	return max
}
