// @program:     tiny-stl
// @file:        copy_test.go.go
// @author:      edte
// @create:      2022-01-17 01:14
// @description:
package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"stl/containers/vector"
)

func TestCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	res := vector.New(vector.WithCap(v.Size()))

	Copy(v.Begin(), v.End(), res.Begin())

	assert.Equal(t, res.String(), "9 3 6 3 1 5 2 3 ")

	fmt.Println(res)
}

func TestCopyN(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	res := vector.New(vector.WithCap(3))
	CopyN(v.Begin(), 3, res.Begin())

	assert.Equal(t, res.String(), "9 3 6 ")

	fmt.Println(res)
}

func TestCopyIf(t *testing.T) {
	op := func(i interface{}) bool {
		return i.(int) > 0
	}

	v := vector.New(vector.WithData([]interface{}{-3, 25, 15, 5, -5, -15}))
	res := vector.New(vector.WithCap(CountIf(v.Begin(), v.End(), op)))

	CopyIf(v.Begin(), v.End(), res.Begin(), op)

	assert.Equal(t, res.String(), "25 15 5 ")

	ForEach(res.Begin(), res.End())
}

func TestCopyBackward(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{-3, 25, 15, 5, -5, -15}))
	res := vector.New(vector.WithCap(v.Size()))

	CopyBackward(v.Begin(), v.End(), res.End())

	fmt.Println(res)
}
