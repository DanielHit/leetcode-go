package array

import (
	"fmt"
	"strconv"
)

// NumMatrix 304. Range Sum Query 2D - Immutable
type NumMatrix struct {
	Array      [][]int
	RegionMap  map[string]int
	PointArray [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{
		Array:      matrix,
		RegionMap:  map[string]int{},
		PointArray: [][]int{},
	}

	pointArray := make([][]int, len(matrix))
	for j := 0; j < len(matrix); j++ {
		pointArray[j] = make([]int, len(matrix[j]))
	}

	// point cal
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 && j == 0 {
				pointArray[i][j] = matrix[0][0]
				continue
			}

			if i == 0 {
				pointArray[0][j] = pointArray[0][j-1] + matrix[0][j]
				continue
			}

			if j == 0 {
				pointArray[i][0] = pointArray[i-1][0] + matrix[i][0]
				continue
			}

			pointArray[i][j] = pointArray[i-1][j] + pointArray[i][j-1] - pointArray[i-1][j-1] + matrix[i][j]

		}
	}
	numMatrix.PointArray = pointArray
	fmt.Println(pointArray)
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	flag := strconv.Itoa(row1) + "_" + strconv.Itoa(col1) + "_" + strconv.Itoa(row2) + "_" + strconv.Itoa(col2)
	if v, ok := this.RegionMap[flag]; ok {
		return v
	}

	s := 0

	s = this.PointArray[row2][col2]
	if row1-1 >= 0 {
		s -= this.PointArray[row1-1][col2]
	}
	if col1-1 >= 0 {
		s -= this.PointArray[row2][col1-1]
	}
	if row1-1 >= 0 && col1-1 >= 0 {
		s += this.PointArray[row1-1][col1-1]
	}

	this.RegionMap[flag] = s
	return s
}
