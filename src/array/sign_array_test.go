package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arraySign(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	is := assert.New(t)
	is.Equal(arraySign(nums), 1)

	nums = []int{-1, -2, 0, -4, 3, 2, 1}
	is.Equal(arraySign(nums), 0)

	nums = []int{-1, 1, -1, 1, -1}
	is.Equal(arraySign(nums), -1)

}
