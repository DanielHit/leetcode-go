package array

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. Remove Duplicates from Sorted List
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	valMap := map[int]bool{}
	p := head
	valMap[p.Val] = true

	for p.Next != nil {
		if valMap[p.Next.Val] {
			p.Next = p.Next.Next
		} else {
			valMap[p.Next.Val] = true
			p = p.Next
		}
	}
	return head
}
