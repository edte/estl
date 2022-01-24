// @program:     tiny-stl
// @file:        queue_test.go.go
// @author:      edte
// @create:      2022-01-15 18:58
// @description:
package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	assert.Equal(t, 1, q.Front())
	assert.Equal(t, 4, q.Back())
	q.Pop()
	assert.Equal(t, 2, q.Front())
	assert.Equal(t, 4, q.Back())
}
