// @program:     tiny-stl
// @file:        sort.go
// @author:      edte
// @create:      2021-12-30 16:07
// @description:
package algorithm

import (
	"stl/comparator"
	"stl/iterator"
)

//  [first,last)
// less comparator
func Sort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewLess()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	for i := first.Clone(); !i.Equal(last); i.Next() {
		for j := i.Clone(); !j.Equal(last); j.Next() {
			if c.Operator(i.Value(), j.Value()) {
				Swap(i, j)
			}
		}
	}
}
