package src

func makeGood(s string) string {
	stack := &Stack{
		Value: []string{},
	}
	for index := 0; index < len(s); index++ {
		if ok, v := stack.get(); ok {
			// 是对应的类型
			if isBadThing(v, string(s[index])) {
				stack.pop()
			} else {
				stack.push(string(s[index]))
			}
		} else {
			stack.push(string(s[index]))
		}
	}

	res := ""
	for _, v := range stack.Value {
		res += v
	}
	return res
}

type Stack struct {
	Value []string
}

// build a stack.
// method push
func (s *Stack) push(a string) {
	s.Value = append(s.Value, a)
}

// build a stack
// method pop
func (s *Stack) pop() (bool, string) {
	if len(s.Value) == 0 {
		return false, ""
	}
	v := s.Value[len(s.Value)-1]
	s.Value = s.Value[:len(s.Value)-1]
	return true, v
}

// method: get the top value of stack
func (s *Stack) get() (bool, string) {
	if len(s.Value) == 0 {
		return false, ""
	}
	return true, s.Value[len(s.Value)-1]
}

// method: check is bad character pair if a - 'A' or 'A' - 'a' true
func isBadThing(a, b string) bool {
	if a[0]-b[0] == 32 || b[0]-a[0] == 32 {
		return true
	}
	return false
}
