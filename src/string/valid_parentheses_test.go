package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	is := assert.New(t)
	is.Equal(longestValidParentheses("()(()"), 2)
	is.Equal(longestValidParentheses("(()"), 2)
	is.Equal(longestValidParentheses(")()()"), 4)
	is.Equal(longestValidParentheses("()(())"), 4)
}

func Test(t *testing.T) {
	fmt.Println(longestValidParentheses("()(())"))
}
