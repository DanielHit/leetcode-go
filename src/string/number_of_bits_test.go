package string

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	num := uint32(00000000000000000000000000001011)
	fmt.Println(hammingWeight(num))
}
