package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	assert.Equal(t, twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
	assert.Equal(t, twoSum([]int{2, 3, 4}, 6), []int{1, 3})
	assert.Equal(t, twoSum([]int{-1, 0}, -1), []int{1, 2})
}
