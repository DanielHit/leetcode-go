package src

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_push(t *testing.T) {
	is := assert.New(t)
	s := &Stack{
		Value: []string{},
	}
	ok, v := s.pop()
	is.False(ok, "")
	is.Equalf(v, "", "")

	s.push("a")
	s.push("b")

	ok, v = s.get()
	is.True(ok)
	is.True(v == "b")

	ok, v = s.pop()
	is.True(ok, "")
	is.True(v == "b")
	ok, v = s.pop()
	is.True(ok)
	is.True(v == "a")
}

func Test_makeGood(t *testing.T) {
	s := makeGood("leEeetcode")
	if s != "leetcode" {
		t.Fail()
	}
	s = makeGood("abBAcC")
	if s != "" {
		t.Fail()
	}
	s = makeGood("s")
	if s != "s" {
		t.Fail()
	}

	s = makeGood("Pp")
	if s != "" {
		t.Fail()
	}

}

func Test_isBadThing(t *testing.T) {
	is := assert.New(t)
	is.True(isBadThing("h", "H"))
	is.True(isBadThing("x", "X"))
	is.False(isBadThing("a", "B"))
	is.False(isBadThing("a", "c"))

	fmt.Println('a' - 'A')
	fmt.Println('A' - 'a')
}
