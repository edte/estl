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

// Copy range of elements
func Copy(first, last iterator.InputIterator, result iterator.OutputIterator) {
	j := result.Clone().(iterator.OutputIterator)
	for i := first.Clone().(iterator.InputIterator); !i.Equal(last); i.Next() {
		j.SetValue(i.Value())
		j.Next()
	}
}

// CopyN copy certain elements of range
// Copies the first n elements from the range beginning at first
// into the range beginning at result.
// The function returns an iterator to the end of the destination
// range (which points to one past the last element copied).
// If n is negative, the function does nothing.
// If the ranges overlap, some of the elements in the range pointed
// by result may have undefined but valid values.
func CopyN(first iterator.InputIterator, n int, result iterator.OutputIterator) iterator.OutputIterator {
	if n < 0 {
		return result
	}

	res := CloneOutput(result)
	i := CloneInput(first)

	for n > 0 {
		res.SetValue(i.Value())
		Advance(i, 1)
		Advance(res, 1)
		n--
	}

	return res
}

// CopyIf return number of elements in range satisfying condition
func CopyIf(first, last iterator.InputIterator, result iterator.OutputIterator, pred functional.Pred) iterator.OutputIterator {
	out := result.Clone().(iterator.OutputIterator)

	for i := first.Clone().(iterator.InputIterator); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			out.SetValue(i.Value())
			out.Next()
		}
	}

	return result
}

// CopyBackward opy range of elements backward
func CopyBackward(first, last, result iterator.BidirectionalIterator) iterator.BidirectionalIterator {
	out := PreB(result)

	for i := PreB(last); i.IsBackEqual(first); i.Pre() {
		out.SetValue(i.Value())
		out.Pre()
	}

	return result
}
