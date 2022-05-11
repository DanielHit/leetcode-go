package src

// problem: leetcode.com/problems/count-sorted-vowel-strings
func countVowelStrings(n int) int {
	if n == 1 {
		return 5
	}
	if n == 2 {
		return 15
	}
	temp := []int{1, 2, 3, 4, 5}
	for i := 3; i <= n; i++ {
		newTemp := []int{}
		sum := 0
		for i2 := range temp {
			sum += temp[i2]
			newTemp = append(newTemp, sum)
		}
		temp = newTemp
	}
	sum := 0
	for _, v := range temp {
		sum += v
	}

	return sum
}
