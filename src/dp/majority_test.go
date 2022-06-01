package dp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majority(t *testing.T) {
	dp, dpChar := majority([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(dp["0_5"] == 4, dpChar["0_5"] == 1)
	assert.True(t, dp["0_5"] == 4)
	assert.True(t, dp["0_3"] == 2)
	assert.True(t, dp["2_3"] == 2)

	dp, dpChar = majority([]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1})
	assert.True(t, dp["3_12"] == 6)
	assert.True(t, dpChar["3_12"] == 2)

}

func Test_majorityII(t *testing.T) {
	dp, dpChar := majorityII([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(dp)
	fmt.Println(dpChar)
	fmt.Println(dp[0][5], dpChar[0][5])
	fmt.Println(dp[0][3], dpChar[0][3])
	fmt.Println(dp[2][3], dpChar[2][3])

	dp, dpChar = majorityII([]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1})
	fmt.Println(dp)
	fmt.Println(dpChar)
	fmt.Println(dp[3][12], dpChar[3][12])
}
