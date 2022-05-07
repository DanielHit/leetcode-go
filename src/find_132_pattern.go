package src

import "fmt"

//
func find132pattern(nums []int) bool {
	stack := &IntStack{Value: nil}
	thridElem := int(^uint(0)>>1) * -1
	fmt.Println(thridElem)

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < thridElem {
			return true
		}
		for {
			if ok, v := stack.get(); ok && v < nums[i] {
				_, thridElem = stack.pop()
			} else {
				break
			}
		}
		stack.push(nums[i])
	}

	return false
}

type IntStack struct {
	Value []int
}

// build a stack.
// method push
func (s *IntStack) push(a int) {
	s.Value = append(s.Value, a)
}

// build a stack
// method pop
func (s *IntStack) pop() (bool, int) {
	if len(s.Value) == 0 {
		return false, 0
	}
	v := s.Value[len(s.Value)-1]
	s.Value = s.Value[:len(s.Value)-1]
	return true, v
}

// method: get the top value of stack
func (s *IntStack) get() (bool, int) {
	if len(s.Value) == 0 {
		return false, 0
	}
	return true, s.Value[len(s.Value)-1]
}
