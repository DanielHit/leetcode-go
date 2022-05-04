package src

import (
	"fmt"
	"testing"
)

func Test_mctFromLeafValues(t *testing.T) {
	arr := []int{6, 2, 4}
	fmt.Println(mctFromLeafValues(arr))
	arr = []int{4, 11}
	fmt.Println(mctFromLeafValues(arr))

	arr = []int{6, 2, 4, 5}
	fmt.Println(mctFromLeafValues(arr))

	arr = []int{15, 13, 5, 3, 15}
	fmt.Println(mctFromLeafValues(arr))

	// 195
	// 270

}
