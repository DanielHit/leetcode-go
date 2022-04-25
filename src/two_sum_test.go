package src

import (
	"testing"
)

// 2022-04-24
func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	if res[0] != 1 || res[1] != 2 {
		t.Fail()
	}

	nums = []int{2, 7, 11, 15}
	target = 9
	res = twoSum(nums, target)
	if res[0] != 0 || res[1] != 1 {
		t.Fail()
	}

	nums = []int{3, 3}
	target = 6
	res = twoSum(nums, target)
	if res[0] != 0 || res[1] != 1 {
		t.Fail()
	}

	nums = []int{2, 5, 5, 11}
	target = 10
	res = twoSum(nums, target)
	if res[0] != 1 || res[1] != 2 {
		t.Fail()
	}

}
