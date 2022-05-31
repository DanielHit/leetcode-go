package dp

import (
	"fmt"
	"testing"
)

func Test_majority(t *testing.T) {
	dp, dpChar := majority([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(dp)
	fmt.Println(dpChar)
	fmt.Println(dp[0][5], dpChar[0][5])
	fmt.Println(dp[0][3], dpChar[0][3])
	fmt.Println(dp[2][3], dpChar[2][3])

	dp, dpChar = majority([]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1})
	fmt.Println(dp)
	fmt.Println(dpChar)
	fmt.Println(dp[3][12], dpChar[3][12])

}
