package src

func mctFromLeafValues(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	if len(arr) == 2 {
		return arr[0] * arr[1]
	}

	l1 := mctFromLeafValues(arr[:len(arr)-1]) + getMaxLeft(arr[:len(arr)-1])*arr[len(arr)-1]
	l2 := mctFromLeafValues(arr[:len(arr)-2]) + arr[len(arr)-1]*arr[len(arr)-2] + getMaxLeft(arr[:len(arr)-2])*getMaxLeft(arr[len(arr)-2:])
	if l1 > l2 {
		return l2
	} else {
		return l1
	}

}

func getMaxLeft(arr []int) int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
