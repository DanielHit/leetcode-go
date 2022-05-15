package src

import (
	"fmt"
	"testing"
)

func Test_numSubarraysWithSum(t *testing.T) {
	nums := []int{1, 0, 1, 0, 1}
	fmt.Println(numSubarraysWithSum(nums, 2))
}
