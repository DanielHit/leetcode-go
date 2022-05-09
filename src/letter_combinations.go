package src

var letterMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {

	res := []string{}
	if len(digits) == 0 {
		return res
	}
	if len(digits) == 1 {
		if v, ok := letterMap[string(digits[0])]; ok {
			for i := 0; i < len(v); i++ {
				res = append(res, string(v[i]))
			}
		}
	}

	leftRes := letterCombinations(digits[:len(digits)-1])
	v := letterMap[string(digits[len(digits)-1])]

	for i := 0; i < len(v); i++ {
		for _, k := range leftRes {
			res = append(res, k+string(v[i]))
		}
	}
	return res
}
