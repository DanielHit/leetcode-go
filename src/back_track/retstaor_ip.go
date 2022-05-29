package back_track

import "strconv"

// 93. Restore IP Addresses
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}

	var res []string
	restoreHelper(s, []string{}, &res, 0)
	return res
}

func restoreHelper(s string, path []string, res *[]string, start int) {
	if start > len(s) || len(path) > 4 {
		return
	}
	if start == len(s) {
		if len(path) == 4 {
			temp := ""
			for i := range path {
				temp += path[i] + "."
			}
			temp = temp[:len(temp)-1]
			*res = append(*res, temp)
			return
		} else {
			return
		}
	}

	// 开始循环
	for i := start; i <= len(s); i++ {
		if validStr(s[start:i]) {
			path = append(path, s[start:i])
			restoreHelper(s, path, res, i)
			path = path[:len(path)-1]
		}
	}

}

func validStr(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	if v > 255 {
		return false
	}
	return true
}
