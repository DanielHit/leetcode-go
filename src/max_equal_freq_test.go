package src

import (
	"testing"
)

func Test_isSubArrayEqualFreq(t *testing.T) {
	nums := []int{2, 2, 1, 1, 5, 3, 3}
	if !isSubArrayEqualFreq(nums) {
		t.Fail()
	}
}

func Test_maxEqualFreq(t *testing.T) {
	nums := []int{2, 2, 1, 1, 5, 3, 3, 5}
	if maxEqualFreq(nums) != 7 {
		t.Fail()
	}
	nums = []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}
	if maxEqualFreq(nums) != 13 {
		t.Fail()
	}
}
