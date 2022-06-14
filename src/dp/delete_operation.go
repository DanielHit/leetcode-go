package dp

// 583. Delete Operation for Two Strings
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	helper(dp, word1, word2)

	return dp[len(word1)][len(word2)]
}

func helper(dp [][]int, w1, w2 string) {
	for i := 0; i <= len(w1); i++ {
		for j := 0; j <= len(w2); j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
}
