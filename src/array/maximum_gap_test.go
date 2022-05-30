package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximumGap(t *testing.T) {
	nums := []int{3, 6, 9, 1}
	assert.Equal(t, maximumGap(nums), 3)

	nums = []int{10}
	assert.Equal(t, maximumGap(nums), 0)

	nums = []int{1, 1, 1, 1}
	assert.Equal(t, maximumGap(nums), 0)

}
