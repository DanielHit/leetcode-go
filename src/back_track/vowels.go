package back_track

import "fmt"

// https://leetcode.com/problems/vowels-of-all-substrings/
var vowelMap = map[string]int{
	"a": 0,
	"e": 0,
	"i": 0,
	"u": 0,
	"o": 0,
}

func countVowels(word string) int64 {
	list := []string{}
	for i := 0; i < len(word); i++ {
		for j := i; j < len(word); j++ {
			subString := word[i : j+1]
			list = append(list, subString)
		}
	}
	fmt.Println(list)

	for _, v := range list {
		count(v)
	}
	fmt.Println(vowelMap)
	sum := int64(0)
	for _, v := range vowelMap {
		sum += int64(v)
	}

	return sum
}
func count(substring string) {
	for i := 0; i < len(substring); i++ {
		if _, ok := vowelMap[string(substring[i])]; ok {
			vowelMap[string(substring[i])]++
		}
	}
}
