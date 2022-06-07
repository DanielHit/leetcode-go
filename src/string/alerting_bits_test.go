package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasAlternatingBits(t *testing.T) {
	assert.True(t, hasAlternatingBits(11) == false)
	assert.True(t, hasAlternatingBits(5) == true)
	assert.True(t, hasAlternatingBits(7) == false)
	assert.True(t, hasAlternatingBits(10) == true)
}
