package src

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	is := assert.New(t)
	is.True(removeDuplicates("deeedbbcccbdaa", 3) == "aa")
	is.True(removeDuplicates("abcd", 2) == "abcd")
	is.True(removeDuplicates("pbbcggttciiippooaais", 2) == "ps")
}

func TestStackDup_Push(t *testing.T) {
	stack := &StackDup{}
	stack.Push("b")
	stack.Push("a")
	stack.Push("a")
	ok, v := stack.Get()
	fmt.Println(ok, v)
	stack.Pop()
	ok, v = stack.Get()
	fmt.Println(ok, v)
}

func TestStackDup_Pop(t *testing.T) {
}
