package test_unit

import (
	"github.com/stretchr/testify/assert"
	utils "go-examples/test_unit"
	"testing"
)

func TestSquare(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(4, utils.Square(2), "他们应该相等")
	assert.NotEqual(4, utils.Square(3))
}
