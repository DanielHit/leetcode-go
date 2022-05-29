package dp

// 爬楼梯. 每次只爬 1个楼梯或者2个楼梯.
// 使用递归
func climb(nums []int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	a := climb(nums[:len(nums)-1])
	b := climb(nums[:len(nums)-2]) + nums[2]

	if b < a {
		return b
	} else {
		return a
	}

}
