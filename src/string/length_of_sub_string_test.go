package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Empty(t *testing.T) {
	assert.Equal(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bbbbb"), 1)
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)
	assert.Equal(t, lengthOfLongestSubstring(" "), 1)
	assert.Equal(t, lengthOfLongestSubstring("au"), 2)
	assert.Equal(t, lengthOfLongestSubstring(""), 0)
	assert.Equal(t, lengthOfLongestSubstring("dvdf"), 3)
}
