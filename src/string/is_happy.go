package string

// 202. Happy Number
func isHappy(n int) bool {
	tempMap := map[int]bool{}
	for n != 1 && !tempMap[n] {
		tempMap[n] = true
		n = getNext(n)
	}
	return n == 1
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n = n / 10
	}
	return sum
}
