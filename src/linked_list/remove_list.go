package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// problem: 203. Remove Linked List Elements
func removeElements(head *ListNode, val int) *ListNode {
	p := head
	for {
		if p != nil && p.Val == val {
			p = p.Next
		} else {
			break
		}
	}

	head = p
	if p == nil {
		return p
	}
	if p.Next == nil {
		if p.Val == val {
			return nil
		} else {
			return head
		}
	}
	q := head.Next

	for {
		if q.Next == nil {
			if q.Val == val {
				p.Next = nil
			}
			break
		}
		if q.Val == val {
			q = q.Next
			p.Next = q
		} else {
			p = q
			q = q.Next
		}
	}
	return head
}
