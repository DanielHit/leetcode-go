package string

func hasAlternatingBits(n int) bool {
	var res []int
	for n > 0 {
		if n%2 == 0 {
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
		if len(res) >= 2 {
			if res[len(res)-1]-res[len(res)-2] == 0 {
				return false
			}
		}
		n = n >> 1
	}
	return true
}
