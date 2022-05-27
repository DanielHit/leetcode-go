package array

import "fmt"

func minDominoRotations(tops []int, bottoms []int) int {
	topMap := map[int]int{}
	bottomMap := map[int]int{}
	for _, v := range tops {
		topMap[v]++
	}
	for _, v := range bottoms {
		bottomMap[v]++
	}

	maxNum := 0
	maxFlag := -1
	for i := 1; i <= 6; i++ {
		if maxNum < topMap[i]+bottomMap[i] {
			maxNum = topMap[i] + bottomMap[i]
			maxFlag = i
		}
	}

	if maxNum < len(tops) {
		return -1
	}
	fmt.Println("the rotation num:", maxFlag, maxNum, topMap, bottomMap)

	count := 0
	if topMap[maxFlag] > bottomMap[maxFlag] {
		// top为主
		for i := 0; i < len(tops); i++ {
			if tops[i] == maxFlag && bottoms[i] == maxFlag {
				continue
			}
			if tops[i] != maxFlag && bottoms[i] != maxFlag {
				return -1
			} else if tops[i] != maxFlag {
				count++
			}
		}
	} else {
		// bottom 为主
		// top为主
		for i := 0; i < len(tops); i++ {
			if tops[i] == maxFlag && bottoms[i] == maxFlag {
				continue
			}
			if tops[i] != maxFlag && bottoms[i] != maxFlag {
				return -1
			} else if bottoms[i] != maxFlag {
				count++
			}
		}
	}
	return count
}
