package array

type Node struct {
	left  *Node
	right *Node
	val   int
}

// problem 164. Maximum Gap
func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	root := &Node{val: nums[0]}

	for i := 1; i < len(nums); i++ {
		insertTree(nums[i], root)
	}

	res := loopTree(root)

	max := 0
	for i := 0; i < len(res)-1; i++ {
		if res[i+1]-res[i] > max {
			max = res[i+1] - res[i]
		}
	}
	return max

}

func loopTree(root *Node) []int {
	if root == nil {
		return []int{}
	}
	if root.left == nil && root.right == nil {
		return []int{root.val}
	}

	var res []int
	res = append(res, loopTree(root.left)...)
	res = append(res, root.val)
	res = append(res, loopTree(root.right)...)
	return res
}

func insertTree(num int, root *Node) {
	if num < root.val {
		if root.left == nil {
			root.left = &Node{val: num}
		} else {
			insertTree(num, root.left)
		}
		return
	}

	if root.right == nil {
		root.right = &Node{val: num}
	} else {
		insertTree(num, root.right)
	}
}
