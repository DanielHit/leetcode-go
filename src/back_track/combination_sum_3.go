package back_track

// problem: https://leetcode.com/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	cand := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	backtracing(cand, &path, &ans, k, n, 0)
	return ans
}

// 动态规划的思路去做
func backtracing(cand []int, path *[]int, res *[][]int, count, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		if len(*path) == count {
			*res = append(*res, *path)
		}
		return
	}
	for i := start; i < len(cand); i++ {
		*path = append(*path, cand[i])
		backtracing(cand, path, res, count, target-cand[i], i+1)
		*path = (*path)[:len(*path)-1]
	}
}
