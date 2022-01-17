// @program:     tiny-stl
// @file:        find.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package algorithm

import (
	"reflect"

	"stl/functional"
	"stl/iterator"
)

// 一些查找算法

// Find value in range
func Find() {

}

// Find element in range
func FindIf() {

}

// Find element in range (negative condition)
func FindIfNot() {

}

// Find last subsequence in range
func FindEnd() {

}

// Find element from set in range
func FindFirstOf() {

}

// Find equal adjacent elements in range
func AdjacentFind() {

}

// Count appearances of value in range
func Count(first, last iterator.InputIterator, val interface{}) int {
	res := 0

	for i := first.Clone(); !i.Equal(last); i.Next() {
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

	for i := first.Clone(); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			res++
		}
	}

	return res
}

// Mismatch return first position where two ranges differ
func Mismatch() {

}

// Equal test whether the elements in two ranges are equal
func Equal(first, last iterator.RandomAccessIterator, first2 iterator.RandomAccessIterator) bool {
	j := first2.Clone()
	for i := first.Clone(); !i.Equal(last); i.Next() {
		if !reflect.DeepEqual(i.Value(), j.Value()) {
			return false
		}
		j.Next()
	}
	return true
}
