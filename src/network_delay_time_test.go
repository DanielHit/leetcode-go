package src

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_networkDelayTime(t *testing.T) {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	is := assert.New(t)
	is.Equal(networkDelayTime(times, 4, 2), 2)

	times = [][]int{{1, 2, 1}}
	is.Equal(networkDelayTime(times, 2, 1), 1)

	times = [][]int{{1, 2, 1}}
	is.Equal(networkDelayTime(times, 2, 2), -1)

	times = [][]int{{1, 2, 1}, {2, 1, 3}}
	fmt.Println(networkDelayTime(times, 2, 2))

}
