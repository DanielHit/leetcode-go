package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findComplement(t *testing.T) {
	assert.Equal(t, findComplement(5), 2)
	assert.Equal(t, findComplement(1), 0)
}
