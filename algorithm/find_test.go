// @program:     tiny-stl
// @file:        find_test.go.go
// @author:      edte
// @create:      2022-01-17 00:57
// @description:
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"stl/functional"
	"testing"

	"stl/containers/vector"
)

func TestFind(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 40}))
	res := Find(v.Begin(), v.End(), 30)
	assert.Equal(t, 30, res.Value())
}

func TestFindNot(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 40}))
	res := Find(v.Begin(), v.End(), 22)
	assert.Equal(t, v.End(), res)
}

func TestFindIf(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 3, 30, 40}))
	res := FindIf(v.Begin(), v.End(), functional.IsOdd)
	assert.Equal(t, 3, res.Value())
}

func TestFindIfNot(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 3, 30, 40}))
	res := FindIfNot(v.Begin(), v.End(), functional.IsOdd)
	assert.Equal(t, 10, res.Value())
}

func TestFindEnd(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}))
	v2 := vector.New(vector.WithData([]interface{}{1, 2, 3}))

	res := FindEnd(v1.Begin(), v1.End(), v2.Begin(), v2.End())

	assert.Equal(t, 5, Position(res))

}

func TestFindFirstOf(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{'a', 'b', 'c', 'A', 'B', 'C'}))
	v2 := vector.New(vector.WithData([]interface{}{'A', 'B', 'C'}))
	res := FindFirstOf(v1.Begin(), v2.End(), v2.Begin(), v2.End())
	assert.Equal(t, 3, Position(res))
}

func TestAdjacentFind(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{5, 20, 5, 30, 30, 20, 10, 10, 20}))
	res := AdjacentFind(v.Begin(), v.End())
	assert.Equal(t, 3, Position(res))
}

func TestCount(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{23, 2, 3, 3, 3, 3, 2, 2}))
	count := Count(v.Begin(), v.End(), 2)
	assert.Equal(t, 3, count)
}

func TestCountIf(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{23, 2, 3, 3, 3, 3, 2, 2}))
	res := CountIf(v.Begin(), v.End(), func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	assert.Equal(t, 3, res)
}

func TestEqual(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{25, 15, 5, -5, -15}))
	v1 := vector.New(vector.WithData([]interface{}{25, 15, 5, -3, -15}))

	equal := Equal(v.Begin(), v.End(), v1.Begin())

	assert.False(t, equal)
}

func TestLowerBound(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 30, 20, 10, 10, 20}))
	Sort(v.Begin(), v.End())
	res := LowerBound(v.Begin(), v.End(), 20)
	assert.Equal(t, 3, Position(res))
}

func TestUpperBound(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 30, 20, 10, 10, 20}))
	Sort(v.Begin(), v.End())
	res := UpperBound(v.Begin(), v.End(), 20)
	assert.Equal(t, 6, Position(res))
}

func TestEqualRange(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 30, 20, 10, 10, 20}))
	Sort(v.Begin(), v.End())
	a, b := EqualRange(v.Begin(), v.End(), 20)

	assert.Equal(t, 3, Position(a))
	assert.Equal(t, 6, Position(b))
}

func TestBinarySearch(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 4, 3, 2, 1}))
	Sort(v.Begin(), v.End())

	assert.True(t, BinarySearch(v.Begin(), v.End(), 3))
	assert.False(t, BinarySearch(v.Begin(), v.End(), 6))
}

func TestSearch(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{10, 20, 30, 40, 50, 60, 70, 80, 90}))
	v2 := vector.New(vector.WithData([]interface{}{40, 50, 60, 70}))
	res := Search(v1.Begin(), v1.End(), v2.Begin(), v2.End())
	assert.Equal(t, 3, Position(res))
}

func TestSearchN(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{10, 20, 30, 30, 20, 10, 10, 20}))
	res := SearchN(v1.Begin(), v1.End(), 2, 30)
	assert.Equal(t, 2, Position(res))
}

func TestMismatch(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{10, 20, 30, 40, 50}))
	v2 := vector.New(vector.WithData([]interface{}{10, 20, 80, 320, 1024}))
	a, b := Mismatch(v1.Begin(), v1.End(), v2.Begin())

	assert.Equal(t, 30, a.Value())
	assert.Equal(t, 80, b.Value())
}
