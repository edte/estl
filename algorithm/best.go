// @program:     tiny-stl
// @file:        BestElement.go
// @author:      edte
// @create:      2021-12-30 16:11
// @description:
package algorithm

import (
	"stl/comparator"
	"stl/iterator"
)

func Max() {

}

func Min() {

}

func MaxMin() {

}

func BestElement(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator, cmp comparator.Comparator) (res iterator.RandomAccessIterator) {
	res = first
	for l := first.Clone(); !l.Equal(last); l.Next() {
		if cmp.Operator(l.Value(), res.Value()) {
			res = l.Clone()
		}
	}
	return res
}

func MinElement(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) (res iterator.RandomAccessIterator) {
	var c comparator.Comparator = comparator.NewLess()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	return BestElement(first, last, c)
}

func MaxElement(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) (res iterator.RandomAccessIterator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	return BestElement(first, last, c)
}

func MaxMinElement() {

}
