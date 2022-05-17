package src

func shortestPathBinaryMatrix(grid [][]int) int {
	return helper_shortest(grid, len(grid)-1, len(grid)-1)
}

func helper_shortest(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		if grid[0][0] == 0 {
			return 1
		} else {
			return -1
		}
	}
	if i < 0 || j < 0 {
		return -1
	}
	if i >= len(grid) || j >= len(grid) {
		return -1
	}
	if i == 0 {
		a := helper_shortest(grid, i, j-1)
		b := helper_shortest(grid, i+1, j-1)
		c := helper_shortest(grid, i+1, j)
		return minAva(minAva(a, b), c)
	}

	a1 := helper_shortest(grid, i-1, j-1)
	a2 := helper_shortest(grid, i-1, j)
	a3 := helper_shortest(grid, i-1, j+1)
	b1 := helper_shortest(grid, i, j-1)
	b2 := helper_shortest(grid, i, j+1)
	c1 := helper_shortest(grid, i+1, j-1)
	c2 := helper_shortest(grid, i+1, j)
	c3 := helper_shortest(grid, i+1, j+1)
	return minAva(minAva(minAva(minAva(minAva(minAva(minAva(a1, a2), a3), b1), b2), c1), c2), c3)
}

func minAva(i, j int) int {
	if i == -1 && j == -1 {
		return -1
	}
	if i == -1 {
		return j
	}
	if j == -1 {
		return i
	}
	if i < j {
		return i
	} else {
		return j
	}
}
