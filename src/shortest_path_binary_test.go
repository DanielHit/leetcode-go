package src

import (
	"fmt"
	"testing"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}

func Test_helper_shortest(t *testing.T) {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(helper_shortest(grid, 2, 2))

}
