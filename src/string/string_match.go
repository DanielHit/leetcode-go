package string

import "sort"
import "strings"

// 1408. String Matching in an Array
func stringMatching(words []string) []string {
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var res []string
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}
