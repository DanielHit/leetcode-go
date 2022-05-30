package array

// problem: 1894. Find the Student that Will Replace the Chalk
func chalkReplacer(chalk []int, k int) int {
	start := -1
	length := len(chalk)
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	k = k % sum
	for {
		if k < 0 {
			return start % length
		}
		start++
		k = k - chalk[start%length]
	}
}
