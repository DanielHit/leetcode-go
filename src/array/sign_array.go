package array

// problem: 1822. Sign of the Product of an Array
func arraySign(nums []int) int {
	negativeCount := 0
	positiveCount := 0
	for _, v := range nums {
		if v == 0 {
			return 0
		} else if v < 0 {
			negativeCount++
		} else {
			positiveCount++
		}
	}
	if negativeCount == 0 {
		return 1
	} else if negativeCount%2 == 0 {
		return 1
	} else {
		return -1
	}
}
