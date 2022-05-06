package src

// problem" https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
func removeDuplicates(s string, k int) string {
	st := &StackDup{Value: nil}

	for i := 0; i < len(s); i++ {
		st.Push(string(s[i]))
		if ok, v := st.Get(); ok && v.Count == k {
			st.Pop()
		}
	}

	res := ""
	for _, node := range st.Value {
		for i := 0; i < node.Count; i++ {
			res += node.Key
		}
	}

	return res
}

type StackNode struct {
	Count int
	Key   string
}

type StackDup struct {
	Value []*StackNode
}

func (statck *StackDup) Push(s string) {
	if statck == nil {
		statck.Value = []*StackNode{}
	}
	if ok, v := statck.Get(); ok {
		if v.Key == s {
			v.Count = v.Count + 1
			return
		}
	}
	statck.Value = append(statck.Value, &StackNode{Count: 1, Key: s})
}

func (statck *StackDup) Pop() (bool, *StackNode) {
	if len(statck.Value) == 0 {
		return true, nil
	}
	statck.Value = statck.Value[:len(statck.Value)-1]
	return true, nil
}

func (statck *StackDup) Get() (bool, *StackNode) {
	if len(statck.Value) <= 0 {
		return false, nil
	}
	return true, statck.Value[len(statck.Value)-1]
}
