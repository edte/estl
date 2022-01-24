// @program:     tiny-stl
// @file:        stack_test.go.go
// @author:      edte
// @create:      2022-01-24 05:06
// @description:
package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	assert.Equal(t, 4, s.Top())
	s.Pop()
	assert.Equal(t, 3, s.Top())
}
