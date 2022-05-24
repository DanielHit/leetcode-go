package string

import "fmt"

// perform
func longestValidParentheses(s string) int {
	stack := &Stack{}
	max := 0

	for i := 0; i < len(s); i++ {
		if ok, c := stack.Peek(); ok {
			if c == "(" && string(s[i]) == ")" {
				stack.Pop()
				if i+1-len(stack.list) > max {
					max = i + 1 - len(stack.list)
				}
				continue
			}
		}

		stack.Push(string(s[i]))
	}

	fmt.Println(stack.list)
	return max

}

// build stack structure
type Stack struct {
	list []string
}

func (s *Stack) Push(c string) {
	s.list = append(s.list, c)
}

func (s *Stack) Pop() (bool, string) {
	if len(s.list) > 0 {
		c := s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
		return true, c
	}
	return false, ""
}

func (s *Stack) Peek() (bool, string) {
	if len(s.list) > 0 {
		return true, s.list[len(s.list)-1]
	}
	return false, ""
}

func (s *Stack) IsEmpty() bool {
	if len(s.list) > 0 {
		return false
	}
	return true
}
