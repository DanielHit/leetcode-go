package rank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge_sort(t *testing.T) {
	nums := []int{4, 5, 2, 7, 9, 3}
	assert.Equal(t, mergeSort(nums), []int{2, 3, 4, 5, 7, 9})

	nums = []int{4, 5, 2}
	assert.Equal(t, mergeSort(nums), []int{2, 4, 5})

	nums = []int{5, 2}
	assert.Equal(t, mergeSort(nums), []int{2, 5})

	nums = []int{2}
	assert.Equal(t, mergeSort(nums), []int{2})

	nums = []int{}
	assert.Equal(t, mergeSort(nums), []int{})

}

func Test_bubble_sort(t *testing.T) {
	nums := []int{4, 5, 2, 7, 9, 3}
	bubbleSort(nums)
	assert.Equal(t, nums, []int{2, 3, 4, 5, 7, 9})

	nums = []int{4, 5, 2}
	bubbleSort(nums)
	assert.Equal(t, nums, []int{2, 4, 5})

	nums = []int{5, 2}
	bubbleSort(nums)
	assert.Equal(t, nums, []int{2, 5})

	nums = []int{2}
	bubbleSort(nums)
	assert.Equal(t, nums, []int{2})

	nums = []int{}
	bubbleSort(nums)
	assert.Equal(t, nums, []int{})

}
