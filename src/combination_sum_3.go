package src

// problem: https://leetcode.com/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	cbDfs(1, k, n, ans, path)
	return ans
}

// 动态规划的思路去做
func cbDfs(num, count, sum int, ans [][]int, path []int) {
	if sum == 0 && count == 0 {
		ans = append(ans, path)
	}
	if sum == 0 || count == 0 {
		return
	}
	if num > 9 {
		return
	}
	if sum >= num && count > 0 {
		path = append(path, num)
		cbDfs(num+1, sum-num, count-1, ans, path)
		path = path[:len(path)-1]
	}

	cbDfs(num+1, sum, count, ans, path)
}
