package dp

import (
	"fmt"
	"testing"
	"time"
)

func Test_climb(t *testing.T) {
	t1 := time.Now()
	nums := []int{1, 100, 1, 1, 100, 1}
	fmt.Println(climb(nums))
	fmt.Println(time.Since(t1))

	l1 := make([]int, 9)
	l1 = append(l1, 2)
	fmt.Println(len(l1))
	fmt.Println(cap(l1))
}
