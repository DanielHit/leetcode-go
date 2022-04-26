package src

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive algorithm
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	var leftPath []string
	var rightPath []string
	if root.Left != nil {
		leftPath = binaryTreePaths(root.Left)
	}

	if root.Right != nil {
		rightPath = binaryTreePaths(root.Right)
	}

	var res []string
	for _, v := range leftPath {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	for _, v := range rightPath {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}

	return res
}
