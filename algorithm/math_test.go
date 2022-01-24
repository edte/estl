// @program:     tiny-stl
// @file:        math_test.go.go
// @author:      edte
// @create:      2021-12-30 16:22
// @description:
package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"stl/comparator"
	"stl/containers/vector"
)

func TestMax(t *testing.T) {
	max := Max(1, 2, 6, 3, 7, 2, 98, -1, 63)
	assert.Equal(t, 98, max)
}

func TestMin(t *testing.T) {
	max := Min(1, 2, 6, 3, 7, 2, 98, -1, 63)
	assert.Equal(t, -1, max)
}

func TestMaxMin(t *testing.T) {
	max, min := MaxMin(1, 2, 6, 3, 7, 2, 98, -1, 63)
	assert.Equal(t, 98, max)
	assert.Equal(t, -1, min)
}

func TestMinElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	e := MinElement(v.Begin(), v.End())
	assert.Equal(t, -1, e.Value())
}

func TestBestElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	j := BestElement(v.Begin(), v.End(), comparator.NewGreater())
	assert.Equal(t, 98, j.Value())
}

func TestMaxMinElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	min, max := MaxMinElement(v.Begin(), v.End())
	assert.Equal(t, 98, max.Value())
	assert.Equal(t, -1, min.Value())
}

func TestBestElementUser(t *testing.T) {
	type user struct {
		age  int
		name string
	}

	v := vector.New(vector.WithData([]interface{}{
		user{
			age:  17,
			name: "a",
		},
		user{
			age:  20,
			name: "b",
		},
		user{
			age:  20,
			name: "a",
		},
		user{
			age:  16,
			name: "c",
		},
		user{
			age:  17,
			name: "c",
		},
		user{
			age:  16,
			name: "e",
		},
	}))

	e := MaxElement(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
		if a.(user).age != b.(user).age {
			return a.(user).age > b.(user).age
		}
		return a.(user).name < b.(user).name
	}))

	assert.Equal(t, "{20 a}", fmt.Sprint(e.Value()))
}

func TestAccumulate(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	res := Accumulate(v.Begin(), v.End())
	assert.Equal(t, 181, res)

	r2 := Accumulate(v.Begin(), v.End(), func(x int, y int) int {
		return x + 3*y
	})

	assert.Equal(t, 543, r2)
}

func TestAdjacentDifference(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4}))
	res := vector.New(vector.WithCapInit(4, 0))
	AdjacentDifference(v.Begin(), v.End(), res.Begin())
	assert.Equal(t, "1 1 1 1 ", res.String())
}

func TestInnerProduct(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{10, 20, 30}))
	v2 := vector.New(vector.WithData([]interface{}{1, 2, 3}))
	res := InnerProduct(v1.Begin(), v1.End(), v2.Begin(), 100)
	assert.Equal(t, 240, res)
}

func TestPartialSum(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{1, 2, 3, 4}))
	res := vector.New(vector.WithCapInit(4, 0))
	PartialSum(v1.Begin(), v1.End(), res.Begin())
	assert.Equal(t, "1 3 6 10 ", res.String())
}

func TestItoa(t *testing.T) {
	v := vector.New(vector.WithCapInit(10, 0))
	Itoa(v.Begin(), v.End(), 100)
	assert.Equal(t, "100 101 102 103 104 105 106 107 108 109 ", v.String())
}
