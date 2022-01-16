// @program:     tiny-stl
// @file:        math_test.go.go
// @author:      edte
// @create:      2021-12-30 16:22
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/comparator"
	"stl/containers/vector"
)

func TestMinElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	fmt.Println(v)

	e := MinElement(v.Begin(), v.End())
	fmt.Println(e.Value())
}

func TestBestElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	fmt.Println(v)
	j := BestElement(v.Begin(), v.End(), comparator.NewGreater())
	fmt.Println(j.Value())
}

func TestMaxMinElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	fmt.Println(MaxMinElement(v.Begin(), v.End()))
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

	fmt.Println(e.Value())
}

func TestAccumulate(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	res := Accumulate(v.Begin(), v.End())
	fmt.Println(res)

	r2 := Accumulate(v.Begin(), v.End(), func(x int, y int) int {
		return x + 3*y
	})

	fmt.Println(r2)
}

func TestAdjacentDifference(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4}))
	res := vector.New(vector.WithCapInit(4, 0))
	AdjacentDifference(v.CBegin(), v.CEnd(), res.OBegin())
	fmt.Println(res)
}

func TestInnerProduct(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{10, 20, 30}))
	v2 := vector.New(vector.WithData([]interface{}{1, 2, 3}))
	fmt.Println(InnerProduct(v1.CBegin(), v1.CEnd(), v2.CBegin(), 100))
}

func TestPartialSum(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{1, 2, 3, 4}))
	res := vector.New(vector.WithCapInit(4, 0))
	PartialSum(v1.CBegin(), v1.CEnd(), res.OBegin())
	fmt.Println(res)
}

func TestItoa(t *testing.T) {
	v := vector.New(vector.WithCapInit(10, 0))
	Itoa(v.Begin(), v.End(), 100)
	fmt.Println(v)
}
