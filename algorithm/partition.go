// @program:     tiny-stl
// @file:        partition.go
// @author:      edte
// @create:      2021-12-30 16:08
// @description:
package algorithm

import (
	"stl/iterator"
)

// IsPartitioned test whether range is partitioned
func IsPartitioned(first, last iterator.RandomAccessIterator, pred func(i interface{}) bool) bool {
	i := first.Clone()

	for !i.Equal(last) {
		if pred(i.Value()) {
			i.Next()
		}
	}

	for !i.Equal(last) {
		if pred(i.Value()) {
			return false
		}
		i.Next()
	}
	return true
}

// Partition range in two
func Partition(first, last iterator.RandomAccessIterator, pred func(i interface{}) bool) iterator.RandomAccessIterator {
	i := first.Clone()
	j := last.Clone().Pre()

	for i.IsFrontEqual(j) {
		if pred(i.Value()) {
			i.Next()
			continue
		}
		if !pred(j.Value()) {
			j.Pre()
			continue
		}
		Swap(i, j)
		i.Next()
		j.Pre()
	}

	return i
}

// StablePartition partition range in two - stable ordering
func StablePartition() {

}

// PartitionCopy partition range into two
func PartitionCopy() {

}

// PartitionPoint get partition point
func PartitionPoint() {

}
