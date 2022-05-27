package rank

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums) / 2
	left := mergeSort(nums[:middle])
	right := mergeSort(nums[middle:])

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

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				swap(nums, i, j)
			}
		}
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
