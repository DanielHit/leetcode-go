package src

// problem: https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/
func mctFromLeafValues(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	if len(arr) == 2 {
		return arr[0] * arr[1]
	}

	n := len(arr)
	maxMap := make([][]int, n)
	keyMap := make([][]int, n)
	for i := 0; i < len(arr); i++ {
		maxMap[i] = make([]int, n)
		keyMap[i] = make([]int, n)
		for j := i; j < len(arr); j++ {
			maxMap[i][j] = -1
			keyMap[i][j] = -1
		}
	}

	ans, _ := dp(0, len(arr)-1, arr, keyMap, maxMap)
	return ans
}

func dp(start, end int, arr []int, keyMap [][]int, maxMap [][]int) (ans int, max int) {
	ans = 2 << 31
	if start > end {
		return 0, 0
	}
	if start == end {
		keyMap[start][end] = 0
		maxMap[start][end] = arr[start]
		return 0, arr[end]
	}

	if keyMap[start][end] >= 0 {
		return keyMap[start][end], maxMap[start][end]
	}

	if start+1 == end {
		keyMap[start][end] = arr[start] * arr[end]
		maxMap[start][end] = maxInt(arr[start], arr[end])
		return arr[start] * arr[end], maxInt(arr[start], arr[end])
	}

	for index := start; index+1 <= end; index++ {
		l, ml := dp(start, index, arr, keyMap, maxMap)
		r, mr := dp(index+1, end, arr, keyMap, maxMap)
		if ans > l+r+ml*mr && ans != 0 {
			ans = l + r + ml*mr
		}
	}
	for i := start; i <= end; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	keyMap[start][end] = ans
	maxMap[start][end] = max

	return ans, max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
