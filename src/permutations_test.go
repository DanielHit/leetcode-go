package src

import (
	"fmt"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permuteUniqueTwo(nums))
}

func Test_permute(t *testing.T) {
	nums := []int{1, 4, 7}
	fmt.Println(permute(nums))

}
