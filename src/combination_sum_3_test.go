package src

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	res := combinationSum3(3, 9)
	fmt.Println(res)
	is := assert.New(t)
	is.Equal(res, [][]int{{1, 2, 4}})

}
