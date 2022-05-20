package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	is := assert.New(t)
	is.Equal(UniquePathsWithObstacles(obstacleGrid), 2)
	obstacleGrid = [][]int{{0, 1}, {0, 0}}
	is.Equal(UniquePathsWithObstacles(obstacleGrid), 1)

	obstacleGrid = [][]int{{0, 0}, {0, 1}}
	is.Equal(UniquePathsWithObstacles(obstacleGrid), 0)

}
