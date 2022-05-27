package rank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge_sort(t *testing.T) {
	nums := []int{4, 5, 2, 7, 9, 3}
	assert.Equal(t, merge_sort(nums), []int{2, 3, 4, 5, 7, 9})

	nums = []int{4, 5, 2}
	assert.Equal(t, merge_sort(nums), []int{2, 4, 5})

	nums = []int{5, 2}
	assert.Equal(t, merge_sort(nums), []int{2, 5})

	nums = []int{2}
	assert.Equal(t, merge_sort(nums), []int{2})

	nums = []int{}
	assert.Equal(t, merge_sort(nums), []int{})

}
