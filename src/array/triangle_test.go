package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8}}
	assert.Equal(t, minimumTotal(triangle), 11)

	triangle = [][]int{{-1}, {2, 3}, {1, -1, -3}}
	assert.Equal(t, minimumTotal(triangle), -1)

	triangle = [][]int{{-10}}
	assert.Equal(t, minimumTotal(triangle), -10)

}

func Test_minimumTotalDp(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	assert.Equal(t, minimumTotalDp(triangle), 11)
}
