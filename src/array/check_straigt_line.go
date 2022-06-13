package array

// 1232. Check If It Is a Straight Line
func checkStraightLine(coordinates [][]int) bool {
	k := float64(0)
	lineFlag := false

	if coordinates[1][0]-coordinates[0][0] == 0 {
		lineFlag = true
	} else {
		k = float64(coordinates[1][1]-coordinates[0][1]) / float64(coordinates[1][0]-coordinates[0][0])
	}

	for i := 1; i < len(coordinates); i++ {
		if lineFlag {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		} else {
			if coordinates[i][0]-coordinates[i-1][0] == 0 {
				return false
			}
			temp := float64(coordinates[i][1]-coordinates[i-1][1]) / float64(coordinates[i][0]-coordinates[i-1][0])
			if temp != k {
				return false
			}
		}

	}

	return true
}
