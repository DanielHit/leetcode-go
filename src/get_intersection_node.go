package src

type ListNode struct {
	Val  int
	Next *ListNode
}

// problem: https://leetcode.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	aCount := 1
	bCount := 1

	pA := headA
	pB := headB

	// count the node length of list a
	for pA.Next != nil {
		pA = pA.Next
		aCount++
	}

	// count the node length of list b
	for pB.Next != nil {
		pB = pB.Next
		bCount++
	}

	// 最后一个节点没有汇聚. 没有公共节点
	if pA != pB {
		return nil
	}

	pA = headA
	pB = headB

	diff := aCount - bCount
	if diff == 0 {
		if pA == pB {
			return pA
		}

		for pA.Next != pB.Next {
			pA = pA.Next
			pB = pB.Next
		}
		return pA.Next

	} else if diff > 0 {
		// a的链路长
		indexP := headA
		for diff > 0 {
			indexP = indexP.Next
			diff--
		}
		for indexP.Next != pB.Next && indexP.Next != nil {
			indexP = indexP.Next
			pB = pB.Next
		}
		return indexP.Next
	} else {
		// b的链路长
		diff = diff * -1
		indexP := headB
		for diff > 0 {
			indexP = indexP.Next
			diff--
		}
		for indexP.Next != pA.Next && indexP.Next != nil {
			indexP = indexP.Next
			pA = pA.Next
		}
		return indexP.Next
	}
}
