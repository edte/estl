// @program:     tiny-stl
// @file:        swap.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package algorithm

import (
	"stl/iterator"
)

func Swap(it1 iterator.RandomAccessIterator, it2 iterator.RandomAccessIterator) {
	t := it1.Value()
	it1.SetValue(it2.Value())
	it2.SetValue(t)
}
