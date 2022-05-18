package string

// problem: 1189. Maximum Number of Balloons
func maxNumberOfBalloons(text string) int {
	countMap := map[string]int{}
	for i := 0; i < len(text); i++ {
		countMap[string(text[i])]++
	}
	bCount := countMap["b"]
	aCount := countMap["a"]
	lCount := countMap["l"] / 2
	oCount := countMap["o"] / 2
	nCount := countMap["n"]

	a := min(min(bCount, aCount), min(lCount, oCount))
	return min(a, nCount)

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
