package array

import (
	"fmt"
	"testing"
)

func Test_minDominoRotations(t *testing.T) {
	bottoms := []int{2, 1, 2, 4, 2, 2}
	tops := []int{5, 2, 6, 2, 3, 2}

	fmt.Println(minDominoRotations(tops, bottoms))

	bottoms = []int{3, 5, 1, 2, 3}
	tops = []int{3, 6, 3, 3, 4}

	fmt.Println(minDominoRotations(tops, bottoms))

}
