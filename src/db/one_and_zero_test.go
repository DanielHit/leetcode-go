package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxForm(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	is := assert.New(t)
	is.Equal(findMaxForm(strs, 5, 3), 4)

	strs = []string{"10", "0", "1"}
	is.Equal(findMaxForm(strs, 1, 1), 2)

	strs = []string{"10", "0001", "111001", "1", "0"}
	is.Equal(findMaxForm(strs, 4, 3), 3)

}

func TestWrong(t *testing.T) {
	//strs := []string{"10", "0001", "111001", "1", "0"}
	//is := assert.New(t)
	//is.Equal(findMaxForm(strs, 4, 3), 3)

	strs := []string{"0", "0", "1", "1"}
	is := assert.New(t)
	is.Equal(findMaxForm(strs, 2, 2), 4)

	strs = []string{"0", "11", "1000", "01", "0", "101", "1", "1", "1", "0", "0", "0", "0", "1", "0", "0110101", "0", "11", "01", "00", "01111", "0011", "1", "1000", "0", "11101", "1", "0", "10", "0111"}
	fmt.Println(findMaxForm(strs, 9, 80))
}

func Test_count(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	zeroMapCount = map[string]int{}
	oneMapCount = map[string]int{}
	for _, v := range strs {
		zeroCount := 0
		oneCount := 0
		for _, c := range v {
			if string(c) == "0" {
				zeroCount++
			} else if string(c) == "1" {
				oneCount++
			}
		}
		zeroMapCount[v] = zeroCount
		oneMapCount[v] = oneCount
	}
	fmt.Println(zeroMapCount)
	fmt.Println(oneMapCount)

	// 核心的使命
}
