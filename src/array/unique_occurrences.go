package array

// 1207. Unique Number of Occurrences
func uniqueOccurrences(arr []int) bool {
	cMap := map[int]int{}
	for _, v := range arr {
		cMap[v]++
	}

	diffMap := map[int]bool{}
	for _, v := range cMap {
		if diffMap[v] {
			return false
		} else {
			diffMap[v] = true
		}
	}

	return true
}
