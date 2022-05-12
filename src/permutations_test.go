package src

import (
	"fmt"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(permuteUnique(nums))
}

func Test_swap(t *testing.T) {
	nums := []int{1, 2}
	swap(nums, 0, 1)
	fmt.Println(nums)
}
