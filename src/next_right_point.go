package src

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Queue struct {
	Value []*Node
}

func (statck *Queue) Push(s *Node) {
	if statck == nil {
		statck.Value = []*Node{}
	}
	if s == nil {
		return
	}
	statck.Value = append(statck.Value, s)
}

func (statck *Queue) Pop() (bool, *Node) {
	if len(statck.Value) == 0 {
		return false, nil
	}
	t := statck.Value[0]
	statck.Value = statck.Value[1:]
	return true, t
}

func (statck *Queue) Get() (bool, *Node) {
	if len(statck.Value) <= 0 {
		return false, nil
	}
	return true, statck.Value[0]
}

func (statck *Queue) IsEmpty() bool {
	if len(statck.Value) <= 0 {
		return true
	}
	return false
}

// 使用queue的效果更好一些
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := &Queue{Value: []*Node{}}
	q.Push(root)

	for !q.IsEmpty() {
		size := len(q.Value)
		for i := 0; i < size; i++ {
			_, t := q.Pop()
			if i != size-1 {
				_, next := q.Get()
				t.Next = next
			}
			if t.Left != nil {
				q.Push(t.Left)
			}
			if t.Right != nil {
				q.Push(t.Right)
			}
		}
	}

	return root
}
