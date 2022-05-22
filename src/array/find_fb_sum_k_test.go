package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinFibonacciNumbers(t *testing.T) {
	is := assert.New(t)
	is.Equal(findMinFibonacciNumbers(7), 2)
	is.Equal(findMinFibonacciNumbers(10), 2)
	is.Equal(findMinFibonacciNumbers(19), 3)

	fmt.Println(findMinFibonacciNumbers(513314))
}
