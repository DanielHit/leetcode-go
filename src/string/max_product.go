package string

import (
	"sort"
)

var totalMap map[string]map[string]bool

// problem 318. Maximum Product of Word Lengths
func maxProduct(words []string) int {
	totalMap = map[string]map[string]bool{}

	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	max := 0
	for i := 0; i < len(words); i++ {
		for j := 1; j < len(words); j++ {
			if len(words[i])*len(words[j]) > max {
				if noCommonLetter(words[i], words[j]) {
					max = len(words[i]) * len(words[j])
				}
			} else {
				break
			}
		}
	}
	return max

}

// 判断是否有
func noCommonLetter(s string, s2 string) bool {
	sMap := map[string]bool{}
	if v, ok := totalMap[s]; ok {
		sMap = v
	} else {
		for _, c := range s {
			sMap[string(c)] = true
		}
		totalMap[s] = sMap
	}

	s2Map := map[string]bool{}
	if v, ok := totalMap[s2]; ok {
		s2Map = v
	} else {
		for _, c := range s2 {
			s2Map[string(c)] = true
		}
		totalMap[s2] = s2Map
	}

	for key := range sMap {
		if s2Map[key] {
			return false
		}
	}

	return true
}
