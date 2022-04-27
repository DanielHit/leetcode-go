package src

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countMatches(t *testing.T) {
	is := assert.New(t)
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"computer", "silver", "phone"}}
	v := countMatches(items, "color", "silver")
	fmt.Println(v)
	is.True(countMatches(items, "color", "silver") == 2, "")

}
