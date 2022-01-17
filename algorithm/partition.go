// @program:     tiny-stl
// @file:        partition.go
// @author:      edte
// @create:      2021-12-30 16:08
// @description:
package algorithm

import (
	"runtime"
	"stl/containers/vector"
	"stl/functional"
	"stl/iterator"
)

// IsPartitioned test whether range is partitioned
func IsPartitioned(first, last iterator.RandomAccessIterator, pred functional.Pred) bool {
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
func Partition(first, last iterator.RandomAccessIterator, pred functional.Pred) iterator.RandomAccessIterator {
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
func StablePartition(first, last iterator.RandomAccessIterator, pred functional.Pred) iterator.RandomAccessIterator {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// 内存足够多，则使用辅助空间
	if m.Frees > 10 {
		return stablePartitionAdaptive(first, last, pred)
	}
	// 原地算法
	return inplaceStablePartition(first, last, pred)
}

func inplaceStablePartition(first, last iterator.RandomAccessIterator, pred functional.Pred) iterator.RandomAccessIterator {
	return nil
}

func stablePartitionAdaptive(first, last iterator.RandomAccessIterator, pred functional.Pred) iterator.RandomAccessIterator {
	v1 := vector.New()
	v2 := vector.New()

	for i := first.Clone(); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			v1.PushBack(i.Value())
		} else {
			v2.PushBack(i.Value())
		}
	}

	//Copy(v1.Begin(), v1.End(),)

	return nil
}

// PartitionCopy partition range into two
func PartitionCopy(first, last iterator.InputIterator, resultTrue, resultFalse iterator.OutputIterator, pred functional.Pred) (iterator.OutputIterator, iterator.OutputIterator) {
	t := resultTrue.Clone()
	f := resultFalse.Clone()

	for i := first.Clone(); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			t.SetValue(i.Value())
			t.Next()
		} else {
			f.SetValue(i.Value())
			f.Next()
		}
	}

	return resultTrue, resultFalse
}

// PartitionPoint get partition point
func PartitionPoint() {

}
