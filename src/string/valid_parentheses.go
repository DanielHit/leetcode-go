package string

import "fmt"

// solve problem 32. Longest Valid Parentheses
func longestValidParentheses(s string) int {
	stack := &Stack{}
	max := 0

	stack.Push(-1)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				if max < i-stack.Peek() {
					max = i - stack.Peek()
				}
			}
		}
	}

	fmt.Println(stack.list)
	return max

}

// build stack structure
type Stack struct {
	list []int
}

func (s *Stack) Push(c int) {
	s.list = append(s.list, c)
}

func (s *Stack) Pop() (bool, int) {
	if len(s.list) > 0 {
		c := s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
		return true, c
	}
	return false, 0
}

func (s *Stack) Peek() int {
	if len(s.list) > 0 {
		return s.list[len(s.list)-1]
	}
	return 0
}

func (s *Stack) IsEmpty() bool {
	if len(s.list) > 0 {
		return false
	}
	return true
}

func (s *Stack) Empty() bool {
	s.list = []int{}
	return true
}
