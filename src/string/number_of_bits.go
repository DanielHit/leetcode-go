package string

import (
	"strconv"
	"strings"
)

func hammingWeight(num uint32) int {
	s := strconv.FormatUint(uint64(num), 2)
	return strings.Count(s, "1")
}
