package tree

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1367. Linked List in Binary Tree
func isSubPath(head *ListNode, root *TreeNode) bool {

	if root == nil {
		return false
	}

	return isPath(head, root) || isPath(head, root.Left) || isPath(head, root.Right)

}

func isPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return isPath(head.Next, root.Left) || isPath(head.Next, root.Right)
}
