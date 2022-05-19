package tree

import "sort"

// problem:783. Minimum Distance Between BST Nodes
func minDiffInBST(root *TreeNode) int {
	res := []int{}
	loopRoot(&res, root)
	sort.Ints(res)

	minDistance := 100000000
	for i := 0; i < len(res)-1; i++ {
		if minDistance > abs(res[i], res[i+1]) {
			minDistance = abs(res[i], res[i+1])
		}
	}
	return minDistance
}

func loopRoot(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	if root.Left != nil {
		loopRoot(res, root.Left)
	}
	if root.Right != nil {
		loopRoot(res, root.Right)
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
