package array

// problem: 747. Largest Number At Least Twice of Others
func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	first := -1
	firstIndex := -1
	second := -1
	secondIndex := -1

	for index, v := range nums {
		if v > second {
			second = v
			secondIndex = index
		}
		if second > first {
			temp := second
			second = first
			first = temp

			tempIndex := secondIndex
			secondIndex = firstIndex
			firstIndex = tempIndex
		}
	}

	if first >= second*2 {
		return firstIndex
	}

	return -1
}
