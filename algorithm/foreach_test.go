// @program:     tiny-stl
// @file:        foreach_test.go.go
// @author:      edte
// @create:      2021-12-30 15:49
// @description:
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"stl/iterator"
	"testing"

	"stl/containers/vector"
)

func TestForEach(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5}))
	ForEach(v.Begin(), v.End())
}

func TestForEach2(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5}))
	res := 0

	ForEach(v.Begin(), v.End(), func(val interface{}) interface{} {
		res += val.(int)
		return ""
	})

	assert.Equal(t, 15, res)

}

func TestForEachIter(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5}))
	ForEachIter(v.Begin(), v.End(), func(i iterator.ForwardIterator) {
		i.SetValue(23)
	})

	assert.Equal(t, 23, v.At(0))
}
