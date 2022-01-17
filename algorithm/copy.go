// @program:     tiny-stl
// @file:        copy.go
// @author:      edte
// @create:      2022-01-11 19:59
// @description:
package algorithm

import (
	"stl/functional"
	"stl/iterator"
)

// CopyIf return number of elements in range satisfying condition
func CopyIf(first, last iterator.InputIterator, result iterator.OutputIterator, pred functional.Pred) iterator.OutputIterator {
	out := result.Clone()

	for i := first.Clone(); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			out.SetValue(i.Value())
			out.Next()
		}
	}

	return result
}

// Copy range of elements
func Copy(first, last iterator.InputIterator, result iterator.RandomAccessIterator) {
	j := result.Clone()
	for i := first.Clone(); !i.Equal(last); i.Next() {
		j.SetValue(i.Value())
		j.Next()
	}
}
