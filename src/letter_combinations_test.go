package src

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	is := assert.New(t)
	res := letterCombinations("23")
	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	is.Equal(res, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
}
