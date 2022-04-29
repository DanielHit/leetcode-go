package src

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// problem: https://leetcode.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	// 开头就类似的
	if headA == headB {
		return headA
	}

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

	fmt.Println("into ")

	pA = headA
	pB = headB

	indexA := &ListNode{}
	indexB := &ListNode{}

	diff := aCount - bCount
	if diff > 0 {
		indexA = headA
		for diff > 0 {
			indexA = indexA.Next
			diff--
		}
		indexB = headB
	} else {
		diff = diff * -1
		indexB = headB
		for diff > 0 {
			indexB = indexB.Next
			diff--
		}
		indexA = headA
	}

	for indexA != indexB && indexA != nil && indexB != nil {
		indexA = indexA.Next
		indexB = indexB.Next
	}
	if indexB == indexA && indexA != nil {
		return indexA
	} else {
		return nil
	}

}
