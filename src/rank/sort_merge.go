package rank

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums) / 2
	left := merge_sort(nums[:middle])
	right := merge_sort(nums[middle:])

	i := 0
	j := 0

	var res []int
	for {
		if i > len(left)-1 || j > len(right)-1 {
			break
		}
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	for ; i <= len(left)-1; i++ {
		res = append(res, left[i])
	}
	for ; j <= len(right)-1; j++ {
		res = append(res, right[j])
	}

	return res
}
