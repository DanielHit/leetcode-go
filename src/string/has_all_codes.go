package string

// problem: 1461. Check If a String Contains All Binary Codes of Size K
func hasAllCodes(s string, k int) bool {
	countMap := map[string]int{}
	for i := 0; i < len(s) && i+k <= len(s); i++ {
		if _, ok := countMap[s[i:i+k]]; ok {

		} else {
			countMap[s[i:i+k]] += 1
		}
	}

	needCount := 1
	for i := 0; i < k; i++ {
		needCount = needCount * 2
	}

	return len(countMap) >= needCount
}
