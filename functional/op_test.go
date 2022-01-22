// @program:     tiny-stl
// @file:        op_test.go.go
// @author:      edte
// @create:      2022-01-22 21:12
// @description:
package functional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNot1(t *testing.T) {
	assert.True(t, IsOdd(1))
	p := Not1(IsOdd)
	assert.False(t, p(1))
}
