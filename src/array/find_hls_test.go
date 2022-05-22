package array

import (
	"fmt"
	"testing"
)

func Test_findLHS(t *testing.T) {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func Test_findLHS_V2(t *testing.T) {
	fmt.Println(findLHS_V2([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func Test_findLHS_V3(t *testing.T) {
	fmt.Println(findLHS_V3([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
