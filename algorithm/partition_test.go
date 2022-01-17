// @program:     tiny-stl
// @file:        partition_test.go.go
// @author:      edte
// @create:      2022-01-14 00:29
// @description:
package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"stl/iterator"
	"testing"

	"stl/containers/vector"
)

func TestIsPartitioned(t *testing.T) {
	cmp := func(i interface{}) bool {
		return i.(int) < 0
	}

	v := vector.New(vector.WithData([]interface{}{9, 4, 3, 6, 3, 1, 5, 2, 99, -1}))

	assert.False(t, IsPartitioned(v.Begin(), v.End(), cmp))

	Partition(v.Begin(), v.End(), cmp)

	assert.True(t, IsPartitioned(v.Begin(), v.End(), cmp))
}

func TestPartition(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 4, 3, 6, 3, 1, 5, 2, 99, -1}))
	Partition(v.Begin(), v.End(), func(i interface{}) bool {
		return i.(int) < 0
	})
	assert.Equal(t, "-1 4 3 6 3 1 5 2 99 9 ", v.String())
}

func TestPartitionCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	isOdd := func(i interface{}) bool {
		return i.(int)%2 != 0
	}

	odd := vector.New(vector.WithCap(CountIf(v.Begin(), v.End(), isOdd)))
	even := vector.New(vector.WithCap(v.Size() - odd.Size()))

	PartitionCopy(v.Begin(), v.End(), odd.Begin(), even.Begin(), isOdd)

	assert.Equal(t, "1 3 5 7 9 ", odd.String())
	assert.Equal(t, "2 4 6 8 ", even.String())
}

type user struct {
	age  int
	name string
}

func getUsers() *vector.Vector {
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

	return v
}

func TestStablePartition(t *testing.T) {
	v := getUsers()

	StablePartition(v.Begin(), v.End(), func(i interface{}) bool {
		u := i.(user)
		return u.age < 18
	})

	assert.Equal(t, "{17 a} {16 c} {17 c} {16 e} {20 b} {20 a} ", v.String())
}

func TestInplaceStablePartition(t *testing.T) {
	v := getUsers()

	inplaceStablePartition(v.Begin(), v.End(), func(i interface{}) bool {
		u := i.(user)
		return u.age < 18
	})

	fmt.Println(v)

	assert.Equal(t, "{17 a} {16 c} {17 c} {16 e} {20 b} {20 a} ", v.String())
}

func TestPartitionPoint(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	odd := vector.New()

	isOdd := func(x interface{}) bool {
		return x.(int)%2 == 1
	}
	StablePartition(v.Begin(), v.End(), isOdd)

	assert.Equal(t, "1 3 5 7 9 2 4 6 8 ", v.String())

	point := PartitionPoint(v.Begin(), v.End(), isOdd)

	assert.Equal(t, "2", fmt.Sprint(point))

	odd.InsertIter(odd.Begin(), v.Begin(), point.(iterator.RandomAccessIterator))

	assert.Equal(t, "1 3 5 7 9 ", fmt.Sprint(odd))

}
