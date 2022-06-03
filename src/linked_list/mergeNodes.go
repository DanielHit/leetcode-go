package linked_list

// 2181. Merge Nodes in Between Zeros
func mergeNodes(head *ListNode) *ListNode {
	var newHead *ListNode
	var p *ListNode
	q := head.Next
	sum := 0
	for q != nil {
		if q.Val == 0 {
			if p == nil {
				p = &ListNode{Val: sum}
				q = q.Next
				newHead = p
				sum = 0
			} else {
				p.Next = &ListNode{Val: sum}
				p = p.Next
				q = q.Next
				sum = 0
			}

		} else {
			sum += q.Val
			q = q.Next
		}
	}
	return newHead
}
