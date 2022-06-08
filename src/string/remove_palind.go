package string

import "math"

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

// 476. Number Complement
func findComplement(num int) int {
	sum := 0
	p := 0
	for num > 0 {
		flag := num % 2
		if flag == 1 {
			flag = 0
		} else {
			flag = 1
		}

		sum += flag * int(math.Pow(2, float64(p)))
		p++
		num = num >> 1
	}
	return sum
}
