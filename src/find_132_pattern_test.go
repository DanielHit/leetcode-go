package src

import (
	"fmt"
	"testing"
)

func Test_find132pattern(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	fmt.Println(find132pattern(nums))
	nums = []int{3, 1, 4, 2}
	fmt.Println(find132pattern(nums))
	nums = []int{-1, 3, 2, 0}
	fmt.Println(find132pattern(nums))
}
