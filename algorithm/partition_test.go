// @program:     tiny-stl
// @file:        partition_test.go.go
// @author:      edte
// @create:      2022-01-14 00:29
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/containers/vector"
)

func TestPartition(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 4, 3, 6, 3, 1, 5, 2, 99, -1}))
	Partition(v.Begin(), v.End(), func(i interface{}) bool {
		return i.(int) < 0
	})
	fmt.Println(v)
}

func TestIsPartitioned(t *testing.T) {
	cmp := func(i interface{}) bool {
		return i.(int) < 0
	}

	v := vector.New(vector.WithData([]interface{}{9, 4, 3, 6, 3, 1, 5, 2, 99, -1}))

	fmt.Println(IsPartitioned(v.Begin(), v.End(), cmp))

	Partition(v.Begin(), v.End(), cmp)

	fmt.Println(IsPartitioned(v.Begin(), v.End(), cmp))
}

func TestPartitionCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	isOdd := func(i interface{}) bool {
		return i.(int)%2 != 0
	}

	odd := vector.New(vector.WithCap(CountIf(v.CBegin(), v.CEnd(), isOdd)))
	even := vector.New(vector.WithCap(v.Size() - odd.Size()))

	PartitionCopy(v.CBegin(), v.CEnd(), odd.OBegin(), even.OBegin(), isOdd)

	fmt.Println(odd)
	fmt.Println(even)
}

func TestStablePartition(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	StablePartition(v.Begin(), v.End(), func(i interface{}) bool {
		return i.(int)%2 == 0
	})
}
