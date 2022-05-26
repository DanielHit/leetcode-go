package string

import (
	"sort"
	"strings"
)

// solve problem:648. Replace Words
func replaceWords(dictionary []string, sentence string) string {
	strList := strings.Split(sentence, " ")
	sort.SliceStable(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	res := []string{}
	for _, str := range strList {
		match := false
		for _, v := range dictionary {
			if count := strings.Index(str, v); count == 0 {
				res = append(res, v)
				match = true
				break
			}
		}
		if match == false {
			res = append(res, str)
		}
	}
	return strings.Join(res, " ")
}
