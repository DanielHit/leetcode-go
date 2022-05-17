package src

func permuteUniqueTwo(nums []int) [][]int {
	var res [][]int
	var path []int
	recursive(&res, path, 0, nums)
	return res
}

func recursive(res *[][]int, path []int, start int, nums []int) {
	if len(path) == 4 {
		*res = append(*res, path)
		return
	} else if start < len(nums) {
		// 先尝试不加元素
		recursive(res, path, start+1, nums)
		// 加了元素的情况
		path = append(path, nums[start])
		recursive(res, path, start+1, nums)
		path = path[:len(path)-1]
	}

}

func permute(nums []int) [][]int {
	var res [][]int
	helper(&res, 0, nums)
	return res
}

func helper(result *[][]int, start int, nums []int) {
	if start == len(nums) {
		var t []int
		for _, i := range nums {
			t = append(t, i)
		}
		*result = append(*result, t)
		return
	}

	// 相应的选择
	for i := start; i < len(nums); i++ {
		swap(i, start, &nums)
		helper(result, start+1, nums)
		swap(start, i, &nums)
	}
}

func swap(i int, start int, nums *[]int) {
	temp := (*nums)[start]
	(*nums)[start] = (*nums)[i]
	(*nums)[i] = temp
}
