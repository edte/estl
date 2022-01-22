// @program:     tiny-stl
// @file:        set_test.go.go
// @author:      edte
// @create:      2022-01-22 17:31
// @description:
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"stl/comparator"
	"stl/containers/vector"
	"testing"
)

func TestMerge(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{1, 3, 4, 6, 10}))
	v2 := vector.New(vector.WithData([]interface{}{2, 5, 7, 8, 9}))
	res := vector.New(vector.WithCap(v1.Size() + v2.Size()))
	Merge(v1.Begin(), v1.End(), v2.Begin(), v2.End(), res.Begin())
	assert.Equal(t, "1 2 3 4 5 6 7 8 9 10 ", res.String())
}

func TestMergeLess(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{1, 3, 4, 6, 10}))
	v2 := vector.New(vector.WithData([]interface{}{2, 5, 7, 8, 9}))

	Sort(v1.Begin(), v1.End(), comparator.NewLess())
	Sort(v2.Begin(), v2.End(), comparator.NewLess())

	res := vector.New(vector.WithCap(v1.Size() + v2.Size()))
	Merge(v1.Begin(), v1.End(), v2.Begin(), v2.End(), res.Begin(), comparator.NewLess())

	assert.Equal(t, "10 9 8 7 6 5 4 3 2 1 ", res.String())
}

func TestInplaceMerge(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{1, 3, 4, 6, 10, 2, 5, 7, 8, 9}))
	InplaceMerge(v1.Begin(), NextN(v1.Begin(), 5), v1.End())
	assert.Equal(t, "1 2 3 4 5 6 7 8 9 10 ", v1.String())
}

func TestInplaceMergeLess(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{10, 6, 4, 3, 1, 9, 8, 7, 5, 2}))
	InplaceMerge(v1.Begin(), NextN(v1.Begin(), 5), v1.End(), comparator.NewLess())
	assert.Equal(t, "10 9 8 7 6 5 4 3 2 1 ", v1.String())
}

func TestIncludes(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{40, 30, 20, 10}))
	v2 := vector.New(vector.WithData([]interface{}{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}))

	Sort(v1.Begin(), v1.End())
	Sort(v2.Begin(), v2.End())

	assert.True(t, Includes(v1.Begin(), v1.End(), v2.Begin(), v2.End()))
}
