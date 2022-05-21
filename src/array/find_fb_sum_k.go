package array

import (
	"fmt"
	"strconv"
)

func findMinFibonacciNumbers(k int) int {
	list := []int{1, 1}

	for i := 2; ; i++ {
		temp := list[i-1] + list[i-2]

		if temp == k {
			return 1
		} else if temp > k {
			break
		}

		list = append(list, temp)
	}

	tempMap := map[string]int{}
	res, _ := helperF(list, &tempMap, k)
	fmt.Println(tempMap)
	return res
}

// 寻找数组
func helperF(list []int, tempMap *map[string]int, k int) (res int, flag bool) {
	n := len(list)

	defer func() {
		if flag {
			(*tempMap)[strconv.Itoa(n)+"_"+strconv.Itoa(k)] = res
		} else {
			(*tempMap)[strconv.Itoa(n)+"_"+strconv.Itoa(k)] = -1
		}
	}()

	if v, ok := (*tempMap)[strconv.Itoa(n)+"_"+strconv.Itoa(k)]; ok {
		if v != -1 {
			return v, true
		} else {
			return -1, false
		}
	}

	if n == 0 && k != 0 {
		return -1, false
	}

	if list[n-1] == k {
		return 1, true
	}

	if k < list[n-1] {
		res, flag = helperF(list[:n-1], tempMap, k)
		return res, flag

	}

	// k > list[n-1]
	a, aFind := helperF(list[:n-1], tempMap, k)

	b, bFind := helperF(list[:n-1], tempMap, k-list[n-1])
	if bFind {
		b = b + 1
	}

	// 两种情况都找到的情况
	if aFind && bFind {
		return min(a, b), true
	}
	// 都没找到
	if !aFind && !bFind {
		return -1, false
	}
	if aFind {
		return a, true
	} else {
		return b, true
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
