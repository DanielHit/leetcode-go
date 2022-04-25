package src

import "testing"
import "github.com/stretchr/testify/assert"

func Test_findMiddleIndex(t *testing.T) {
	// test find middle index
	is := assert.New(t)
	nums := []int{2, 3, -1, 8, 4}
	is.True(findMiddleIndex(nums) == 3)
	nums = []int{1, -1, 4}
	is.True(findMiddleIndex(nums) == 2)

	nums = []int{2, 3}
	is.True(findMiddleIndex(nums) == -1)
}
