// @program:     tiny-stl
// @file:        foreach.go
// @author:      edte
// @create:      2021-12-30 15:36
// @description:
package algorithm

import (
	"stl/functional"
	"stl/iterator"
)

func ForEach(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator, f functional.Op) {
	i := first.Clone()
	for ; !i.Equal(last); i.Next() {
		f(i.Value())
	}
}
