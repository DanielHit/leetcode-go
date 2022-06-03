package tree

import "sync"

// 105. Construct Binary Tree from Preorder and Inorder Traversal
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}

	rootVal := preorder[0]
	index := 0
	for index, _ = range inorder {
		if inorder[index] == rootVal {
			break
		}
		index++
	}

	wc := sync.WaitGroup{}
	wc.Add(2)
	root := &TreeNode{Val: rootVal}

	go func() {
		defer wc.Done()
		root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	}()

	go func() {
		defer wc.Done()
		root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	}()

	wc.Wait()
	return root
}
