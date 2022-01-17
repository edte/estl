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
	i := Clone(first)

	for !i.Equal(last) && pred(i.Value()) {
		i.Next()
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
	i := Clone(first)
	j := Pre(last)

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

// TODO：快慢指针失败了，并不稳定，修改为归并递归 O(log^n)
func inplaceStablePartition(first, last iterator.RandomAccessIterator, pred functional.Pred) iterator.RandomAccessIterator {
	i := Clone(first)
	j := Clone(first)

	for !i.Equal(last) && !j.Equal(last) {
		if pred(i.Value()) {
			i.Next()
			continue
		}
		j = Clone(i)
		for !j.Equal(last) && !pred(j.Value()) {
			j.Next()
		}
		if j.Equal(last) {
			break
		}
		Swap(i, j)
		i.Next()
		j.Next()
	}

	return first
}

func stablePartitionAdaptive(first, last iterator.RandomAccessIterator, pred functional.Pred) iterator.RandomAccessIterator {
	v1 := vector.New()
	v2 := vector.New()

	ForEach(first, last, func(val interface{}) {
		if pred(val) {
			v1.PushBack(val)
		} else {
			v2.PushBack(val)
		}
	})

	Copy(v1.Begin(), v1.End(), first)
	Copy(v2.Begin(), v2.End(), NextN(first, Distance(v1.Begin(), v1.End())))

	return nil
}

// PartitionCopy partition range into two
func PartitionCopy(first, last iterator.InputIterator, resultTrue, resultFalse iterator.OutputIterator, pred functional.Pred) (iterator.OutputIterator, iterator.OutputIterator) {
	t := CloneOutput(resultTrue)
	f := CloneOutput(resultFalse)

	ForEach(first, last, func(val interface{}) {
		if pred(val) {
			t.SetValue(val)
			t.Next()
		} else {
			f.SetValue(val)
			f.Next()
		}

	})

	return resultTrue, resultFalse
}

// PartitionPoint get partition point
// Returns an iterator to the first element in the partitioned
//range [first,last) for which pred is not true, indicating its
//partition point.
//The elements in the range shall already be partitioned, as if
//partition had been called with the same arguments.
func PartitionPoint(first, last iterator.ForwardIterator, pred functional.Pred) iterator.ForwardIterator {
	n := Distance(first, last)

	for n > 0 {
		it := CloneForward(first)
		step := n / 2
		Advance(it, step)

		if pred(it.Value()) {
			first = Next(it)
			n -= step + 1
		} else {
			n = step
		}
	}

	return first
}
