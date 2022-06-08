package string

// 1332. Remove Palindromic Subsequences
func removePalindromeSub(s string) int {
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}

	return 1

}
