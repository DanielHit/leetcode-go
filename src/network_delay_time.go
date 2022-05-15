package src

import "strconv"

// problem:743
func networkDelayTime(times [][]int, n int, k int) int {
	// find map to contain reverse path
	pathMap := map[int]int{}
	lengthMap := map[string]int{}
	alreadyMap := map[string]bool{}
	for _, pair := range times {
		u := pair[0]
		v := pair[1]
		w := pair[2]
		if alreadyMap[strconv.Itoa(v)+"_"+strconv.Itoa(u)] {
			continue
		}

		pathMap[v] = u
		lengthMap[strconv.Itoa(v)+"_"+strconv.Itoa(u)] = w

	}

	// reverse find
	start := n
	costTime := 0
	for {
		if v, ok := pathMap[start]; ok {
			costTime += lengthMap[strconv.Itoa(start)+"_"+strconv.Itoa(v)]
			start = v
			if v == k {
				break
			}
		} else {
			break
		}
	}
	if start == k {
		return costTime
	} else {
		return -1
	}

}
