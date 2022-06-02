package dp

// 1035. Uncrossed Lines
// time complexity o(m*n)
func maxUncrossedLines(nums1 []int, nums2 []int) int {

	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 0; i <= len(nums1); i++ {
		for j := 0; j <= len(nums2); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}

			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(nums1)][len(nums2)]

}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
