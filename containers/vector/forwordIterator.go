// @program:     tiny-stl
// @file:        forwordIterator.go
// @author:      edte
// @create:      2022-01-17 19:04
// @description:
package vector

import (
	"fmt"
	"stl/iterator"
)

type ForwardIterator struct {
	InputIterator
	OutputIterator
	data *Vector
	pos  int
}

func NewForwardIterator(data *Vector, pos int) *ForwardIterator {
	return &ForwardIterator{data: data, pos: pos}
}

func (iter *ForwardIterator) Next() iterator.ForwardIterator {
	if iter.pos == iter.data.Size() {
		return iter
	}
	iter.pos++
	return iter
}

func (iter *ForwardIterator) Clone() iterator.ForwardIterator {
	return NewForwardIterator(iter.data, iter.pos)
}

func (iter *ForwardIterator) HasNext() bool {
	return iter.pos == iter.data.Size()
}

func (iter *ForwardIterator) Equal(other iterator.ForwardIterator) bool {
	i, ok := other.(*ForwardIterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos == iter.pos
}

func (iter *ForwardIterator) IsValid() bool {
	return iter.pos >= -1 && iter.pos <= iter.data.Size()
}

func (iter *ForwardIterator) Value() interface{} {
	return iter.data.At(iter.pos)
}

func (iter *ForwardIterator) SetValue(value interface{}) {
	iter.data.SetAt(iter.pos, value)
}

func (iter *ForwardIterator) String() string {
	return fmt.Sprint(iter.data.At(iter.pos))
}
