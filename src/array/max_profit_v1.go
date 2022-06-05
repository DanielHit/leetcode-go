package array

// 121. Best Time to Buy and Sell Stock
func maxProfitV1(prices []int) int {
	min := 1000001
	diff := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else {
			if v-min > diff {
				diff = v - min
			}
		}
	}
	return diff
}

// 122. Best Time to Buy and Sell Stock II
func maxProfitV2(prices []int) int {
	sum := 0
	min := 1000000
	for _, v := range prices {
		if v < min {
			min = v
		} else {
			sum += v - min
			min = v
		}
	}
	return sum
}
