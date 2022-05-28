package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	assert.Equal(t, findMedianSortedArrays(nums1, nums2), 2.5)
	assert.Equal(t, findMedianSortedArrays([]int{1, 3}, []int{2}), float64(2))
}

func Test_findMedianSortedArrays1(t *testing.T) {

}
