// @program:     tiny-stl
// @file:        foreach.go
// @author:      edte
// @create:      2021-12-30 15:36
// @description:
package algorithm

import (
	"fmt"
	"stl/functional"
	"stl/iterator"
)

// ForEach Apply function to range
// Applies function fn to each of the elements in the range [first,last).
func ForEach(first, last iterator.InputIterator, ops ...functional.Op) {
	op := func(val interface{}) {
		fmt.Print(val)
		fmt.Print(" ")
	}
	if len(ops) != 0 {
		op = ops[0]
	}

	for i := CloneInput(first); !i.Equal(last); i.Next() {
		op(i.Value())
	}
}

func ForEachW(first, last iterator.ForwardIterator, f functional.IterOp) {
	for i := CloneForward(first); !i.Equal(last); i.Next() {
		f(i.(iterator.ForwardIterator))
	}
}
