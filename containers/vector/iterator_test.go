// @program:     tiny-stl
// @file:        iterator_test.go.go
// @author:      edte
// @create:      2022-01-15 17:19
// @description:
package vector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"stl/iterator"
	"testing"
)

func TestIterator_String(t *testing.T) {
	v := New(WithData([]interface{}{1, 2, 3, 4}))
	fmt.Println(v.Begin().Next())
}

func TestVector_String(t *testing.T) {
	v := New(WithData([]interface{}{1, 2, 3, 4}))
	i := NewIterator(v, 0)
	fmt.Println(i.String())
}

func TestIterator_IsBackEqual(t *testing.T) {
	v := New(WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	i1 := v.Begin()
	i2 := i1.Clone().(iterator.RandomAccessIterator).NextN(3)

	assert.True(t, i1.IsFrontEqual(i2))
	assert.True(t, i1.IsFront(i2))
	assert.False(t, i1.IsBackEqual(i2))
	assert.False(t, i1.IsBack(i2))
}
