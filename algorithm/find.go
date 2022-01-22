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
func FindEnd(first1, last1, first2, last2 iterator.ForwardIterator, pred ...functional.BinaryPred) iterator.ForwardIterator {
	p := functional.EqualTo
	if len(pred) != 0 {
		p = pred[0]
	}

	if first2.Equal(last2) {
		return last1
	}
	res := CloneBidirectional(last1)

	i := CloneBidirectional(first1)
	j := CloneBidirectional(first2)

	for !i.Equal(last1) {
		i1 := CloneBidirectional(i)
		j1 := CloneBidirectional(j)

		for !i1.Equal(last1) && !j1.Equal(last2) && p(i1.Value(), j1.Value()) {
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
func FindFirstOf(first1, last1 iterator.InputIterator, first2, last2 iterator.ForwardIterator, pred ...functional.BinaryPred) iterator.InputIterator {
	p := functional.EqualTo
	if len(pred) != 0 {
		p = pred[0]
	}

	i := CloneInput(first1)

	for !i.Equal(last1) {
		for j := CloneForward(first2); !j.Equal(last2); j.Next() {
			if p(i.Value(), j.Value()) {
				return i
			}
		}
		i.Next()
	}
	return last1
}

// AdjacentFind find equal adjacent elements in range
func AdjacentFind() {

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

// Mismatch return first position where two ranges differ
func Mismatch(first1, last1, first2 iterator.InputIterator, pred functional.BinaryPred) (iterator.InputIterator, iterator.InputIterator) {

	return nil, nil
}

// LowerBound return iterator to lower bound
func LowerBound(first, last iterator.ForwardIterator, val int, cmp ...comparator.Comparator) iterator.ForwardIterator {

	return nil
}

// UpperBound return iterator to upper bound
func UpperBound(first, last iterator.ForwardIterator, val int, cmp ...comparator.Comparator) iterator.ForwardIterator {

	return nil
}

// EqualRange get subrange of equal elements
func EqualRange(first, last iterator.ForwardIterator, val int, cmp ...comparator.Comparator) (iterator.ForwardIterator, iterator.ForwardIterator) {

	return nil, nil
}

// BinarySearch test if value exists in sorted sequence
func BinarySearch(first, last iterator.ForwardIterator, val int, cmp ...comparator.Comparator) bool {

	return false
}

// Search range for subsequence
func Search(first1, last1, first2, last2 iterator.ForwardIterator, pred functional.BinaryPred) iterator.ForwardIterator {

	return nil
}

// SearchN search range for elements
func SearchN(first, last iterator.ForwardIterator, count, val int, pred functional.BinaryPred) iterator.ForwardIterator {

	return nil
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
