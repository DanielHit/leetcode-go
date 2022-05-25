package array

import (
	"fmt"
	"testing"
)

func Test_jump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))

	nums = []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums))

}
