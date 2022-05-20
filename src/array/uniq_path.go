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
