package array

// findMedianSortedArrays: solve problem  4. Median of Two Sorted Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mergeArray []int
	firstIndex := 0
	secondIndex := 0
	for {
		if firstIndex > len(nums1)-1 || secondIndex > len(nums2)-1 {
			break
		}
		if nums1[firstIndex] < nums2[secondIndex] {
			mergeArray = append(mergeArray, nums1[firstIndex])
			firstIndex++
		} else {
			mergeArray = append(mergeArray, nums2[secondIndex])
			secondIndex++
		}
	}
	for ; firstIndex < len(nums1); firstIndex++ {
		mergeArray = append(mergeArray, nums1[firstIndex])
	}
	for ; secondIndex < len(nums2); secondIndex++ {
		mergeArray = append(mergeArray, nums2[secondIndex])
	}
	middle := (len(nums1) + len(nums2)) / 2
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(mergeArray[middle])
	} else {
		return float64(mergeArray[middle]+mergeArray[middle-1]) / 2
	}
}
