// @program:     tiny-stl
// @file:        outputIterator.go
// @author:      edte
// @create:      2022-01-16 21:01
// @description:
package vector

import (
	"fmt"

	"stl/iterator"
)

type OutputIterator struct {
	data *Vector
	pos  int
}

func NewOutputIterator(data *Vector, pos int) *OutputIterator {
	return &OutputIterator{data: data, pos: pos}
}

func (iter *OutputIterator) Next() iterator.OutputIterator {
	if iter.pos == iter.data.Size() {
		return iter
	}
	iter.pos++
	return iter
}

func (iter *OutputIterator) Clone() iterator.OutputIterator {
	return NewOutputIterator(iter.data, iter.pos)
}

func (iter *OutputIterator) HasNext() bool {
	return iter.pos == iter.data.Size()
}

func (iter *OutputIterator) Equal(other iterator.OutputIterator) bool {
	i, ok := other.(*OutputIterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos == iter.pos
}

func (iter *OutputIterator) IsValid() bool {
	return iter.pos >= -1 && iter.pos <= iter.data.Size()
}

func (iter *OutputIterator) SetValue(value interface{}) {
	iter.data.SetAt(iter.pos, value)
}

func (iter *OutputIterator) String() string {
	return fmt.Sprint(iter.data.At(iter.pos))
}
