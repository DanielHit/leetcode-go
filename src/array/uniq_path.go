package array

import "strconv"

// UniquePathsWithObstacles solve problem:63. Unique Paths II
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	pathMap := map[string]int{}

	return helper(obstacleGrid, &pathMap, m-1, n-1)
}

func helper(grid [][]int, m *map[string]int, p, q int) int {
	if (*m)[strconv.Itoa(p)+"_"+strconv.Itoa(q)] > 0 {
		return (*m)[strconv.Itoa(p)+"_"+strconv.Itoa(q)]
	}

	if q == 0 {
		flag := true
		for j := 0; j <= p; j++ {
			if grid[j][0] == 1 {
				flag = false
				break
			}
		}
		if flag {
			(*m)[strconv.Itoa(p)+"_"+strconv.Itoa(q)] = 1
			return 1
		} else {
			return 0
		}
	}
	if p == 0 {
		flag := true
		for i := 0; i <= q; i++ {
			if grid[0][i] == 1 {
				flag = false
				break
			}
		}
		if flag {
			(*m)[strconv.Itoa(p)+"_"+strconv.Itoa(q)] = 1
			return 1
		} else {
			return 0
		}
	}

	sum := 0
	if grid[p][q] == 1 {
		return sum
	}
	if grid[p-1][q] == 0 {
		sum += helper(grid, m, p-1, q)
	}
	if grid[p][q-1] == 0 {
		sum += helper(grid, m, p, q-1)
	}

	(*m)[strconv.Itoa(p)+"_"+strconv.Itoa(q)] = sum
	return sum
}

// dp的算法再来解决一遍. 思考下解决方案
//func helper_2(grid [][]int, res *[][]int) {
//	if grid[0][0] == 1 {
//		(*res)[0][0] = 0
//	} else {
//		(*res)[0][0] = 1
//	}
//
//	for j := 1; j < len(grid[0]); j++ {
//		if grid[0][j] == 1 {
//			res[0][j] = 0
//		} else {
//			if
//		}
//	}
//
//	for i := 0; i < len(grid[0]); i++ {
//
//	}
//
//	for i := 1; i < len(grid); i++ {
//		for j := 1; j < len(grid[0]); j++ {
//			if grid[i][j] == 1 {
//				(*res)[i][j] = 0
//			} else {
//				if ((*res)[i-1][j] > (*res)[i][j-1]) && ((*res)[i-1][j] != 0) {
//					(*res)[i][j] = (*res)[i][j-1] + 1
//				} else if ((*res)[i][j-1]) != 0 {
//					(*res)[i][j] = (*res)[i-1][j] + 1
//				} else {
//					(*res)[i][j] = 0
//				}
//			}
//		}
//	}
//}
