// @program:     tiny-stl
// @file:        permutation_test.go.go
// @author:      edte
// @create:      2022-01-24 21:53
// @description:
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"stl/comparator"
	"stl/containers/vector"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 3, 5, 4, 2}))
	NextPermutation(v.Begin(), v.End())
	assert.Equal(t, "1 4 2 3 5 ", v.String())

	v = vector.New(vector.WithData([]interface{}{1, 2, 3}))
	NextPermutation(v.Begin(), v.End())
	assert.Equal(t, "1 3 2 ", v.String())
	NextPermutation(v.Begin(), v.End())
	assert.Equal(t, "2 1 3 ", v.String())
	NextPermutation(v.Begin(), v.End())
	assert.Equal(t, "2 3 1 ", v.String())
	NextPermutation(v.Begin(), v.End())
	assert.Equal(t, "3 1 2 ", v.String())
	NextPermutation(v.Begin(), v.End())
	assert.Equal(t, "3 2 1 ", v.String())
	NextPermutation(v.Begin(), v.End())
	assert.Equal(t, "1 2 3 ", v.String())

	NextPermutation(v.Begin(), v.End(), comparator.NewLess())
	assert.Equal(t, "3 2 1 ", v.String())
	NextPermutation(v.Begin(), v.End(), comparator.NewLess())
	assert.Equal(t, "3 1 2 ", v.String())
	NextPermutation(v.Begin(), v.End(), comparator.NewLess())
	assert.Equal(t, "2 3 1 ", v.String())
	NextPermutation(v.Begin(), v.End(), comparator.NewLess())
	assert.Equal(t, "2 1 3 ", v.String())
	NextPermutation(v.Begin(), v.End(), comparator.NewLess())
	assert.Equal(t, "1 3 2 ", v.String())
	NextPermutation(v.Begin(), v.End(), comparator.NewLess())
	assert.Equal(t, "1 2 3 ", v.String())
}

func TestPrePermutation(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3}))

	PrePermutation(v.Begin(), v.End())
	assert.Equal(t, "3 2 1 ", v.String())
	PrePermutation(v.Begin(), v.End())
	assert.Equal(t, "3 1 2 ", v.String())
	PrePermutation(v.Begin(), v.End())
	assert.Equal(t, "2 3 1 ", v.String())
	PrePermutation(v.Begin(), v.End())
	assert.Equal(t, "2 1 3 ", v.String())
	PrePermutation(v.Begin(), v.End())
	assert.Equal(t, "1 3 2 ", v.String())
	PrePermutation(v.Begin(), v.End())
	assert.Equal(t, "1 2 3 ", v.String())
}
