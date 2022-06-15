package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subMaxArray(t *testing.T) {
	arr := []int{-1, 2, 4, -10, 8, 9}
	assert.Equal(t, subMaxArray(arr), 17)
}
