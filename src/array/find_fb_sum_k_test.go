package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinFibonacciNumbers(t *testing.T) {
	is := assert.New(t)
	is.Equal(findMinFibonacciNumbers(7), 2)
	is.Equal(findMinFibonacciNumbers(10), 2)
	is.Equal(findMinFibonacciNumbers(19), 3)
}
