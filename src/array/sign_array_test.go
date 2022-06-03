package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_arraySign(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	is := assert.New(t)
	is.Equal(arraySign(nums), 1)

	nums = []int{-1, -2, 0, -4, 3, 2, 1}
	is.Equal(arraySign(nums), 0)

	nums = []int{-1, 1, -1, 1, -1}
	is.Equal(arraySign(nums), -1)

}

func TestArray(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	fmt.Println(preorder[0:1])
	fmt.Println(preorder[1:])
	fmt.Println(preorder[1:2])
	fmt.Println(preorder[1:3])
	wc := sync.WaitGroup{}
	wc.Add(2)
	go func() {
	}()
}
