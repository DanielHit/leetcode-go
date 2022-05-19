package back_track

import "fmt"

var vowelMap = map[string]int{
	"a": 0,
	"e": 0,
	"i": 0,
	"u": 0,
	"o": 0,
}

// https://leetcode.com/problems/vowels-of-all-substrings/
func countVowels(word string) int64 {
	res := []string{}
	helper(&res, 0, "", word)
	count := int64(0)
	fmt.Println(vowelMap)
	fmt.Println(res)
	for _, v := range vowelMap {
		count += int64(v)
	}
	return count
}

func countOfVowels(s string) {
	for i := 0; i < len(s); i++ {
		vowelMap[string(s[i])]++
	}
}

func helper(res *[]string, start int, substring, word string) {

	if start >= len(word) {
		*res = append(*res, substring)
		countOfVowels(substring)
		return
	}

	helper(res, start+1, substring, word)
	for i := start; i < len(word); i++ {
		substring = word[i : start+1]
		helper(res, start+1, substring, word)
		substring = word[i:start]
	}
}
