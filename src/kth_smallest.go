package src

import "github/leetcode-go/src/tree"

// problem: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *tree.TreeNode, k int) int {
	resList := searchTree(root)
	return resList[k-1]
}

func searchTree(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, searchTree(root.Left)...)
	res = append(res, root.Val)
	res = append(res, searchTree(root.Right)...)
	return res
}
