package back_track

func SplitIpString(s string) [][]string {
	var res [][]string
	helperSplitIp(s, 0, &res, []string{})
	return res
}

func helperSplitIp(s string, start int, result *[][]string, path []string) {
	if start == len(s) {
		*result = append(*result, path)
		return
	}
	for i := start + 1; i < len(s); i++ {
		helperSplitIp(s, i, result, path)

		// ip的特殊的定制
		path = append(path, string(s[i]))
		helperSplitIp(s, i+1, result, path)
		path = path[:len(path)-1]
	}

}
