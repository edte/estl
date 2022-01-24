// @program:     tiny-stl
// @file:        find.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package algorithm

import (
	"stl/comparator"
	"stl/functional"
	"stl/iterator"
)

// 一些查找算法

// Find value in range
// Returns an iterator to the first element in the range [first,last)
// that compares equal to val. If no such element is found, the
// function returns last.
// The function uses operator== to compare the individual elements
// to val.
func Find(first, last iterator.InputIterator, val interface{}) (res iterator.InputIterator) {
	for i := CloneInput(first); !i.Equal(last); i.Next() {
		if i.Value() == val {
			return i
		}

	}
	return last
}

// FindIf find element in range
func FindIf(first, last iterator.InputIterator, pred functional.Pred) iterator.InputIterator {
	for i := CloneInput(first); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			return i
		}
	}
	return last
}

// FindIfNot find element in range (negative condition)
func FindIfNot(first, last iterator.InputIterator, pred functional.Pred) iterator.InputIterator {
	return FindIf(first, last, functional.Not1(pred))
}

// FindEnd find last subsequence in range
func FindEnd(first1, last1, first2, last2 iterator.ForwardIterator, preds ...functional.BinaryPred) iterator.ForwardIterator {
	pred := functional.DefaultBinaryPared(preds...)
	if first2.Equal(last2) {
		return last1
	}
	res := CloneBidirectional(last1)

	i := CloneBidirectional(first1)
	j := CloneBidirectional(first2)

	for !i.Equal(last1) {
		i1 := CloneBidirectional(i)
		j1 := CloneBidirectional(j)

		for !i1.Equal(last1) && !j1.Equal(last2) && pred(i1.Value(), j1.Value()) {
			i1.Next()
			j1.Next()

			if j1.Equal(last2) {
				res = CloneBidirectional(i)
				break
			}

			if i1.Equal(last1) {
				return res
			}
		}

		i.Next()
	}

	return res
}

// FindFirstOf find element from set in range
func FindFirstOf(first1, last1 iterator.InputIterator, first2, last2 iterator.ForwardIterator, preds ...functional.BinaryPred) iterator.InputIterator {
	pred := functional.DefaultBinaryPared(preds...)
	i := CloneInput(first1)

	for !i.Equal(last1) {
		for j := CloneForward(first2); !j.Equal(last2); j.Next() {
			if pred(i.Value(), j.Value()) {
				return i
			}
		}
		i.Next()
	}
	return last1
}

// AdjacentFind find equal adjacent elements in range
func AdjacentFind(first, last iterator.ForwardIterator, preds ...functional.BinaryPred) iterator.ForwardIterator {
	pred := functional.DefaultBinaryPared(preds...)

	if first.Equal(last) {
		return last
	}

	i := CloneForward(first)
	j := NextForward(first)

	for !i.Equal(j) {
		if pred(i.Value(), j.Value()) {
			return i
		}
		i.Next()
		j.Next()
	}

	return last
}

// Count appearances of value in range
func Count(first, last iterator.InputIterator, val interface{}) int {
	res := 0

	for i := CloneInput(first); !i.Equal(last); i.Next() {
		if i.Value() == val {
			res++
		}
	}

	return res
}

// CountIf return number of elements in range satisfying condition
// returns the number of elements in the range [first,last) for which pred is true.
func CountIf(first, last iterator.InputIterator, pred functional.Pred) int {
	res := 0

	for i := CloneInput(first); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			res++
		}
	}

	return res
}

// LowerBound return iterator to lower bound
func LowerBound(first, last iterator.ForwardIterator, val interface{}, cmp ...comparator.Comparator) iterator.ForwardIterator {
	var c comparator.Comparator = comparator.NewLess()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	distance := Distance(first, last)

	i := CloneForward(first)

	for distance > 0 {
		it := CloneForward(i)
		step := distance / 2
		Advance(it, step)

		if c.Operator(it.Value(), val) {
			i = NextForward(it)
			distance -= step + 1
			continue
		}

		distance = step
	}

	return i
}

// UpperBound return iterator to upper bound
func UpperBound(first, last iterator.ForwardIterator, val interface{}, cmp ...comparator.Comparator) iterator.ForwardIterator {
	var c = comparator.Reserve(comparator.NewLess())
	if len(cmp) != 0 {
		c = cmp[0]
	}

	distance := Distance(first, last)

	i := CloneForward(first)

	for distance > 0 {
		it := CloneForward(i)
		step := distance / 2
		Advance(it, step)

		if c.Operator(val, it.Value()) {
			i = NextForward(it)
			distance -= step + 1
			continue
		}

		distance = step
	}

	return i
}

// EqualRange get subrange of equal elements
func EqualRange(first, last iterator.ForwardIterator, val interface{}, cmp ...comparator.Comparator) (iterator.ForwardIterator, iterator.ForwardIterator) {
	it := LowerBound(first, last, val, cmp...)

	return it, UpperBound(it, last, val, cmp...)
}

// BinarySearch test if value exists in sorted sequence
func BinarySearch(first, last iterator.ForwardIterator, val int, cmp ...comparator.Comparator) bool {
	it := LowerBound(first, last, val, cmp...)
	return !it.Equal(last) && functional.EqualTo(val, it.Value())
}

// Search range for subsequence
func Search(first1, last1, first2, last2 iterator.ForwardIterator, preds ...functional.BinaryPred) iterator.ForwardIterator {
	if first2.Equal(last2) {
		return first1
	}

	pred := functional.DefaultBinaryPared(preds...)

	i1 := CloneForward(first1)
	i2 := CloneForward(first2)

	for !i1.Equal(i2) {
		it1 := CloneForward(i1)
		it2 := CloneForward(i2)

		for pred(it1.Value(), it2.Value()) {
			it1.Next()
			it2.Next()
			if it2.Equal(last2) {
				return i1
			}
			if it1.Equal(last1) {
				return last1
			}
		}
		i1.Next()
	}

	return last1
}

// SearchN search range for elements
// searches the range [first,last) for a sequence of count
// elements, each comparing equal to val (or for which pred returns
// true).
func SearchN(first, last iterator.ForwardIterator, count, val int, preds ...functional.BinaryPred) iterator.ForwardIterator {
	pred := functional.DefaultBinaryPared(preds...)

	limit := NextN(first, Distance(first, last)-count)

	for i := CloneForward(first); !i.Equal(limit); i.Next() {
		it := CloneForward(i)
		j := 0

		for pred(it.Value(), val) {
			it.Next()
			j++
			if j == count {
				return i
			}
		}
	}

	return last
}

// Mismatch return first position where two ranges differ
// compares the elements in the range [first1,last1) with those
// in the range beginning at first2, and returns the first element
// of both sequences that does not match.
func Mismatch(first1, last1, first2 iterator.InputIterator, preds ...functional.BinaryPred) (iterator.InputIterator, iterator.InputIterator) {
	pred := functional.DefaultBinaryPared(preds...)

	i := CloneInput(first1)
	j := CloneInput(first2)

	for !i.Equal(last1) && pred(i.Value(), j.Value()) {
		i.Next()
		j.Next()
	}

	return i, j
}

// Equal test whether the elements in two ranges are equal
func Equal(first, last iterator.Iterator, first2 iterator.Iterator) bool {
	j := Clone(first2)
	for i := Clone(first); !i.Equal(last); i.Next() {
		if !i.Equal(j) {
			return false
		}
		j.Next()
	}
	return true
}
