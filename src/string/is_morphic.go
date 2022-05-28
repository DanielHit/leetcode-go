package string

// isIsomorphic:205. Isomorphic Strings
func isIsomorphic(s string, t string) bool {
	matchMap := map[string]string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if v, ok := matchMap[string(s[i])]; ok {
			str += v
		} else {
			str += string(t[i])
			matchMap[string(s[i])] = string(t[i])
		}
	}

	if str != t {
		return false
	}

	matchMap = map[string]string{}
	str = ""
	for i := 0; i < len(t); i++ {
		if v, ok := matchMap[string(t[i])]; ok {
			str += v
		} else {
			str += string(s[i])
			matchMap[string(t[i])] = string(s[i])
		}
	}

	return str == s

}
