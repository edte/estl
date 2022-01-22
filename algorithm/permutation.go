// @program:     tiny-stl
// @file:        permutation.go
// @author:      edte
// @create:      2021-12-30 16:10
// @description:
package algorithm

import (
	"stl/comparator"
	"stl/functional"
	"stl/iterator"
)

// 一些排序组合算法

// LexicographicalCompare lexicographical less-than comparison
// returns true if the range [first1,last1) compares
// lexicographically less than the range [first2,last2).
func LexicographicalCompare(first1, last1, first2, last2 iterator.InputIterator, cmp ...comparator.Comparator) bool {
	c := comparator.Default(cmp...)

	i := CloneInput(first1)
	j := CloneInput(first2)

	for !i.Equal(last1) {
		if j.Equal(last2) || c.Operator(j.Value(), i.Value()) {
			return false
		} else if c.Operator(i.Value(), i.Value()) {
			return true
		}
		i.Next()
		j.Next()
	}

	return !j.Equal(last2)
}

// NextPermutation transform range to next permutation
// rearranges the elements in the range [first,last) into the next
// lexicographically greater permutation.
func NextPermutation(first, last iterator.BidirectionalIterator, cmp ...comparator.Comparator) bool {
	//cmp := comparator.Default(cmp...)

	return true
}

// PrePermutation transform range to previous permutation
func PrePermutation(first, last iterator.BidirectionalIterator, cmp ...comparator.Comparator) bool {
	//cmp := comparator.Default(cmp...)

	return true
}

// IsPermutation test whether range is permutation of another
func IsPermutation(first1, last1, first2 iterator.InputIterator, pred functional.BinaryPred) bool {
	return false
}
