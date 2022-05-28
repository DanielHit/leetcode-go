package array

//  containsNearbyDuplicate: solve 219. Contains Duplicate II
func containsNearbyDuplicate(nums []int, k int) bool {
	temp := map[int]int{}
	for index, v := range nums {
		if lastIndex, ok := temp[v]; ok {
			if abs(index, lastIndex) <= k {
				return true
			} else {
				temp[v] = index
			}
		} else {
			temp[v] = index
		}
	}
	return false
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
