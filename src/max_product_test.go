package src

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_theMaxOddLengthPalindromic(t *testing.T) {
	s := "abad"
	is := assert.New(t)
	res := theMaxOddLengthPalindromic(s)
	fmt.Println(res)
	is.True(res == 3)
	s = "abcd12321"
	res = theMaxOddLengthPalindromic(s)
	fmt.Println(res)
	is.True(res == 5)
	s = "ababbb"
	is.True(theMaxOddLengthPalindromic(s) == 3, 3)
}

func Test_maxProduct(t *testing.T) {
	s := "ababbb"
	is := assert.New(t)
	is.True(maxProduct(s) == 9)
	s = "zaaaxbbby"
	is.True(maxProduct(s) == 9)
}
