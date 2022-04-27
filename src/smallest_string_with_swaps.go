package src

// https://leetcode.com/problems/count-items-matching-a-rule/
func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	typeCountMap := map[string]int{}
	colorCountMap := map[string]int{}
	nameCountMap := map[string]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < len(items); j++ {
			if i == 0 {
				if _, ok := typeCountMap[items[j][i]]; ok {
					typeCountMap[items[j][i]]++
				} else {
					typeCountMap[items[j][i]] = 1
				}
			}
			if i == 1 {
				if _, ok := colorCountMap[items[j][i]]; ok {
					colorCountMap[items[j][i]]++
				} else {
					colorCountMap[items[j][i]] = 1
				}
			}
			if i == 2 {
				if _, ok := nameCountMap[items[j][i]]; ok {
					nameCountMap[items[j][i]]++
				} else {
					nameCountMap[items[j][i]] = 1
				}
			}
		}
	}

	switch ruleKey {
	case "type":
		return typeCountMap[ruleValue]
	case "color":
		return colorCountMap[ruleValue]
	case "name":
		return nameCountMap[ruleValue]
	}

	return 0
}
