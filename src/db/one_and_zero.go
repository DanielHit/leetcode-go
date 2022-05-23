package db

var zeroMapCount map[string]int
var oneMapCount map[string]int

func findMaxForm(strs []string, m int, n int) int {
	ans := 0
	initZeroAndOneMap(strs)
	helper(strs, &ans, 0, m, n, []string{})
	return ans
}

func initZeroAndOneMap(strs []string) {
	zeroMapCount = map[string]int{}
	oneMapCount = map[string]int{}
	for _, v := range strs {
		zeroCount := 0
		oneCount := 0
		for _, c := range v {
			if string(c) == "0" {
				zeroCount++
			} else if string(c) == "1" {
				oneCount++
			}
		}
		zeroMapCount[v] = zeroCount
		oneMapCount[v] = oneCount
	}
}

func helper(strs []string, ans *int, start int, m, n int, subString []string) {
	if m < 0 || n < 0 {
		return
	}

	if start >= len(strs) {
		if len(subString) >= *ans {
			*ans = len(subString)
		}
		return
	}

	helper(strs, ans, start+1, m, n, subString)
	subString = append(subString, strs[start])
	zeroM := zeroMapCount[strs[start]]
	oneN := oneMapCount[strs[start]]
	helper(strs, ans, start+1, m-zeroM, n-oneN, subString)
	subString = subString[:len(subString)-1]
}
