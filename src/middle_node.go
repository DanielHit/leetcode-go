package src

// solution: https://leetcode.com/problems/middle-of-the-linked-list/submissions/
func middleNode(head *ListNode) *ListNode {
	count := 0
	p := head
	for p != nil {
		count++
		p = p.Next
	}

	middleIndex := count/2 + 1
	p = head
	for i := 1; i < middleIndex; i++ {
		p = p.Next
	}
	return p
}
